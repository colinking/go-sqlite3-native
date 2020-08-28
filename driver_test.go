package sqlite3native

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/segmentio/events/v2"
	_ "github.com/segmentio/events/v2/sigevents"
	"github.com/segmentio/events/v2/text"
	"github.com/stretchr/testify/require"
)

func init() {
	// Note: we disable the prefix here for tests, to reduce events's verbosity
	events.DefaultHandler = &text.Handler{
		Output: os.Stdout,
	}
}

func TestDriverE2E(tt *testing.T) {
	for _, test := range []struct {
		name    string
		setup   string           // sql to run on the DB via sqlite3
		sql     string           // sql to run
		results [][]driver.Value // expected results
	}{
		{
			name: "simple multi-row integer select all",
			setup: `
				PRAGMA journal_mode=WAL;
				CREATE TABLE table1 (column1 int);
				INSERT INTO table1 (column1) VALUES (123);
				INSERT INTO table1 (column1) VALUES (456);
			`,
			sql: "select * from table1;",
			results: [][]driver.Value{
				{int64(123)},
				{int64(456)},
			},
		},
	} {
		tt.Run(test.name, func(t *testing.T) {
			require := require.New(t)

			// This test will be run on a sqlite3 DB in a temporary directory.
			dir, err := ioutil.TempDir("", "go-sqlite3-native-*")
			require.NoError(err)
			defer func() {
				// require.NoError(os.RemoveAll(dir))
			}()

			// Write the setup SQL to a file so we can pipe it into sqlite3
			inputPath := filepath.Join(dir, "input.sql")
			err = ioutil.WriteFile(inputPath, []byte(test.setup), 0644)
			require.NoError(err)

			// Execute the setup SQL on this temporary SQLite DB:
			dbPath := filepath.Join(dir, "test.db")
			events.Log("test path: %s", dbPath)
			sh := fmt.Sprintf("cat %s | sqlite3 %s", inputPath, dbPath)
			cmd := exec.Command("bash", "-c", sh)
			stdout, err := cmd.Output()
			require.NoError(err)
			events.Log("%s\nstdout: %s", sh, stdout)

			// Open this SQLite DB with our Go client:
			db, err := sql.Open("sqlite3-native", dbPath)
			require.NoError(err)
			defer func() {
				require.NoError(db.Close())
			}()

			// Prepare the test SQL command
			stmt, err := db.PrepareContext(context.Background(), test.sql)
			require.NoError(err)

			// Run the test SQL command:
			rows, err := stmt.QueryContext(context.Background())
			require.NoError(err)

			// Verify we got the expected results.
			cols, err := rows.Columns()
			require.NoError(err)
			results := [][]driver.Value{}
			for rows.Next() {
				// To scan in an arbitrary length of items, we need to do some
				// pointer juggling:
				row := make([]driver.Value, len(cols))
				ptrs := make([]interface{}, len(cols))
				for i := range cols {
					ptrs[i] = &row[i]
				}
				err := rows.Scan(ptrs...)
				require.NoError(err)
				results = append(results, row)
			}
			require.NoError(rows.Err())
			require.Equal(test.results, results)
		})
	}
}
