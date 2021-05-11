package cmd

import (
	"fmt"

	"github.com/VKCOM/noverify/src/linter"
)

func Check(ctx *AppContext) (int, error) {
	config := ctx.MainConfig.linter.Config()

	bindConfigValuesWithFlags(ctx, config)

	ruleSets, err := parseExternalRules(ctx.ParsedFlags.RulesList)
	if err != nil {
		return 1, fmt.Errorf("preload external rules: %v", err)
	}

	for _, rset := range ruleSets {
		config.Checkers.DeclareRules(rset)
	}

	ctx.MainConfig.rulesSets = append(ctx.MainConfig.rulesSets, ruleSets...)

	if ctx.ParsedFlags.DisableCache {
		config.CacheDir = ""
	}

	if ctx.MainConfig.AfterFlagParse != nil {
		ctx.MainConfig.AfterFlagParse(InitEnvironment{
			RuleSets: ctx.MainConfig.rulesSets,
			MetaInfo: ctx.MainConfig.linter.MetaInfo(),
		})
	}

	return mainNoExit(ctx)
}

func bindConfigValuesWithFlags(ctx *AppContext, config *linter.Config) {
	config.CheckAutoGenerated = ctx.ParsedFlags.CheckAutoGenerated
	config.CheckAutoGenerated = ctx.ParsedFlags.Debug
	config.DebugParseDuration = ctx.ParsedFlags.DebugParseDuration
	config.MaxConcurrency = ctx.ParsedFlags.MaxConcurrency
	config.StubsDir = ctx.ParsedFlags.StubsDir
	config.CacheDir = ctx.ParsedFlags.CacheDir
	config.IgnoreTriggerError = ctx.ParsedFlags.IgnoreTriggerError
	config.ConservativeBaseline = ctx.ParsedFlags.ConservativeBaseline
	config.ApplyQuickFixes = ctx.ParsedFlags.ApplyQuickFixes
	config.KPHP = ctx.ParsedFlags.KPHP
}
