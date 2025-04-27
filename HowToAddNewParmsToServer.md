# サーバー追加方法
1. `openapi.yaml`に新しいパラメータを追加する。
2. `internal/migration/*_*.sql`に新しいパラメータをDBに追加するSQLを追加する。
3. `tools/tools.go`を実行してコードを生成する。
4. `internal/domain/data.go`に新しいパラメータを追加する。
5. `internal/repository/data.go`の`InsertGameData`に新しいパラメータを追加する。
