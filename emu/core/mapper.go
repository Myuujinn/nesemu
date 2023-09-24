package core

const (
	NoMapper           = 0
	MMC1               = 1
	Unrom              = 2
	Cnrom              = 3
	Mmc3               = 4
	Mmc5               = 5
	Ffef4              = 6
	Aorom              = 7
	Ffef3              = 8
	Mmc2               = 9
	Mmc4               = 10
	ColorDreams        = 11
	Ffef6              = 12
	HundredIn1         = 15
	Bandai             = 16
	Ffef8              = 17
	Jaleco             = 18
	Namcot             = 19
	NintendoDiskSystem = 20
	KonamiVrc4a        = 21
	KonamiVrc2a        = 22
	KonamiVrc2a2       = 23
	KonamiVrc6         = 24
	KonamiVrc4b        = 25
	IremG101           = 32
	TaitoTC            = 33
	RomSwitch          = 34
	Rambo1             = 64
	IremH3001          = 65
	Gnrom              = 66
	Sunsoft3           = 67
	Sunsoft4           = 68
	Sunsoft5           = 69
	Camerica           = 71
	Irem74             = 78
	Pirate             = 91
)

func MapperName(mapper byte) string {
	switch mapper {
	case NoMapper:
		return "None"
	case MMC1:
		return "MMC1"
	case Unrom:
		return "UNROM"
	case Cnrom:
		return "CNROM"
	case Mmc3:
		return "MMC3"
	case Mmc5:
		return "MMC5"
	case Ffef4:
		return "FFE F4"
	case Aorom:
		return "AOROM"
	case Ffef3:
		return "FFE F3"
	case Mmc2:
		return "MMC2"
	case Mmc4:
		return "MMC4"
	case ColorDreams:
		return "Color Dreams"
	case Ffef6:
		return "FFE F6"
	case HundredIn1:
		return "100 In 1"
	case Bandai:
		return "Bandai"
	case Ffef8:
		return "FFE F8"
	case Jaleco:
		return "Jaleco"
	case Namcot:
		return "Namcot"
	case NintendoDiskSystem:
		return "Nintendo Disk System"
	case KonamiVrc4a:
		return "Konami VRC4a"
	case KonamiVrc2a:
		return "Konami VRC2a"
	case KonamiVrc2a2:
		return "Konami VRC2a2"
	case KonamiVrc6:
		return "Konami VRC6"
	case KonamiVrc4b:
		return "Konami VRC4b"
	case IremG101:
		return "Irem G-101"
	case TaitoTC:
		return "Taito TC"
	case RomSwitch:
		return "ROM Switch"
	case Rambo1:
		return "Rambo 1"
	case IremH3001:
		return "Irem H-3001"
	case Gnrom:
		return "GNROM"
	case Sunsoft3:
		return "Sunsoft 3"
	case Sunsoft4:
		return "Sunsoft 4"
	case Sunsoft5:
		return "Sunsoft 5"
	case Camerica:
		return "Camerica"
	case Irem74:
		return "Irem 74"
	case Pirate:
		return "Pirate"
	default:
		return "Unknown Mapper"
	}
}
