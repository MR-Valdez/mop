package fire

import (
	"github.com/wowsims/mop/sim/core"
	"github.com/wowsims/mop/sim/mage"
)

func (fire *FireMage) registerHotfixes() {
	// 2025-11-13 - Critical Mass value lowered back to the 5.4.8 value of 1.3x (was 1.5x)
	//fire.criticalMassMultiplier += 0.2

	// 2025-11-13 - Pyroblast's damage increase to Pyroblast lowered to 15%. (was 30%)
	fire.AddStaticMod(core.SpellModConfig{
		ClassMask:  mage.MageSpellPyroblast,
		Kind:       core.SpellMod_DamageDone_Pct,
		FloatValue: 0.15,
	})

	// 2025-07-01 - Combustion Ignite scaling increased to 50% (was 20%).
	fire.combustionDotDamageMultiplier += 0.3
}
