package feature

import (
	"gin-web/app/common/helpers"
)

const (
	All = "All"
	Android = "Android"
	IPhone = "iPhone"
	WeChat = "WeChat"
)

type Expression struct {
	Operator string
	Origin	 string
}

type Condition struct {
	OS			string		`yaml:"os"`
	Version		Expression		`yaml:"version"`
}

type Feature struct {
	Name		string		`yaml:"name"`
	Condition	Condition	`yaml:"condition"`
}

type Features struct {
	features []Feature
	myFeatures []string
}

func (f *Features) Init(YAMLFile string, env Condition) {
	helper.YAML(YAMLFile, &f.features)
	for _, feature := range f.features  {
		if feature.Condition.OS != All && feature.Condition.OS != env.OS{
			continue
		}
		if feature.Condition.Version.Operator == "" {
			if feature.Condition.Version.Origin != "" {
				feature.Condition.Version.Operator = "="
			} else {
				feature.Condition.Version.Operator = ">"
			}
		}
		ok := helper.VersionCompare(
			env.Version.Origin,
			feature.Condition.Version.Origin,
			feature.Condition.Version.Operator)
		if !ok {
			continue
		}
		f.myFeatures = append(f.myFeatures, feature.Name)
	}
}

func (f *Features) Has(name string) bool {
	for _, feature := range f.myFeatures {
		if feature == name {
			return true
		}
	}
	return false
}