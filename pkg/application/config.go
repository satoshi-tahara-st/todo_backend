package application

//go:generate enumer -transform snake -json -yaml -type=Stage -trimprefix Stage
type Stage int

const (
	//ローカル
	StageLocal Stage = iota
	//検証環境
	StageStaging
	//本番環境
	StageProduction
)

// // Database DBの設定
// type Database struct {
// 	Driver string
// 	DBName string
// 	Ref    ConnectionSetting `yaml:"ref"`
// 	Upd    ConnectionSetting `yaml:"upd`
// }

type Config struct {
}
