package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug})))
	a := map[string]any{"a": 1, "b": "c"}
	aj, _ := json.MarshalIndent(a, "", "\t")

	//log.Printf("qua xin: a: %s", aj)
	// Output:
	// qua xin: a: {
	//       "a": 1,
	//       "b": "c"
	// }

	_ = aj
	slog.Info("kho doc vai lin", "a", aj)
	// Output:
	// INFO kho doc vai lin a="{\n\t\"a\": 1,\n\t\"b\": \"c\"\n}"
}
