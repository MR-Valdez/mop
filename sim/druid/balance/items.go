package balance

import (
	"time"

	"github.com/wowsims/mop/sim/core"
	"github.com/wowsims/mop/sim/core/stats"
	"github.com/wowsims/mop/sim/druid"
)

func init() {
}

// T15 Balance
var ItemSetRegaliaOfTheHauntedForest = core.NewItemSet(core.ItemSet{
	Name:                    "Regalia of the Haunted Forest",
	DisabledInChallengeMode: true,
	Bonuses: map[int32]core.ApplySetBonus{
		2: func(_ core.Agent, setBonusAura *core.Aura) {
			// Increases the critical strike chance of Starsurge by 10%.
			setBonusAura.AttachSpellMod(core.SpellModConfig{
				Kind:       core.SpellMod_BonusCrit_Percent,
				ClassMask:  druid.DruidSpellStarsurge,
				FloatValue: 10,
			})
		},
		4: func(agent core.Agent, setBonusAura *core.Aura) {
			// Nature's Grace now also grants 1000 critical strike and 1000 mastery for its duration.
			moonkin := agent.(*BalanceDruid)
			statValue := 1000.0
			aura := moonkin.NewTemporaryStatsAura(
				"Nature's Grace 4pT15",
				core.ActionID{SpellID: 138351},
				stats.Stats{stats.HasteRating: statValue, stats.CritRating: statValue, stats.MasteryRating: statValue},
				time.Second*15,
			)

			moonkin.AddEclipseCallback(func(_ Eclipse, gained bool, sim *core.Simulation) {
				if gained {
					aura.Activate(sim)
				}
			})
		},
	},
})
