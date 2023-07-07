package gamedata

import (
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/roboden-game/serverapi"
)

func CalcDifficultyScore(config serverapi.ReplayLevelConfig, pointsAllocated int) int {
	score := 100

	switch config.RawGameMode {
	case "reverse":
		score -= (config.BossDifficulty - 2) * 20
		score += (3 - config.CreepDifficulty) * 15
		score += (config.DronesPower - 1) * 15
		score += config.StartingResources * 5
		score -= (config.InitialCreeps - 1) * 15
		score -= (config.TechProgressRate - 5) * 5
		score += (config.Teleporters - 1) * 5
		score += (config.OilRegenRate - 2) * 5
		score += (config.Resources - 2) * 15

	case "classic":
		if config.InterfaceMode < 2 {
			score += 5
		}
		if config.NumCreepBases != 0 {
			score += (config.CreepDifficulty - 3) * 15
			if config.SuperCreeps {
				score += 45
			}
		} else {
			score += (config.CreepDifficulty - 3) * 10
			if config.SuperCreeps {
				score += 35
			}
		}
		score += 10 - (config.Resources * 5)
		score += (config.NumCreepBases - 2) * 15
		score += (config.BossDifficulty - 1) * 15
		score += (config.CreepSpawnRate - 1) * 10
		score += (config.InitialCreeps - 1) * 10
		score -= config.StartingResources * 4
		score += 5 - (config.Teleporters * 5)
		score += 10 - (config.OilRegenRate * 5)
		score += 20 - pointsAllocated
	case "arena", "inf_arena":
		if config.InterfaceMode == 0 {
			score += 5
		}
		if config.RawGameMode == "inf_arena" {
			score += (config.ArenaProgression - 1) * 15
			score += 30 - (config.Resources * 15)
			score += config.InitialCreeps * 5
		} else {
			score += (config.ArenaProgression - 1) * 10
			score += 20 - (config.Resources * 10)
			score += config.InitialCreeps * 10
		}
		score += (config.CreepDifficulty - 3) * 20
		score -= config.StartingResources * 2
		score += 5 - (config.Teleporters * 5)
		score += 10 - (config.OilRegenRate * 5)
		score += 20 - pointsAllocated
	}

	if config.FogOfWar {
		score += 5
	}

	return gmath.ClampMin(score, 1)
}
