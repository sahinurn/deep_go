package main

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	familyOffset = 3
	gunOffset    = 4
	houseOffset  = 5
	hpOffset     = 12
	mpOffset     = 22
)

const (
	statsBaseMask = 0xF
	hpMask        = uint32(0x3FF << hpOffset)
	manaMask      = uint32(0x3FF << mpOffset)
)

type GamePerson struct {
	stats  uint16
	name   [42]byte
	info   uint32
	xCoord int32
	yCoord int32
	zCoord int32
	gold   uint32
}

func NewGamePerson(options ...Option) GamePerson {
	p := GamePerson{}

	for _, o := range options {
		o(&p)
	}

	return p
}

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.xCoord = int32(x)
		person.yCoord = int32(y)
		person.zCoord = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		mask := uint32(mana) << mpOffset
		person.info = person.info | mask
	}
}

func (p *GamePerson) Mana() int {
	val := p.info & manaMask
	return int(val >> mpOffset)
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		mask := uint32(health) << hpOffset
		person.info = person.info | mask
	}
}

func (p *GamePerson) Health() int {
	val := p.info & hpMask
	return int(val >> hpOffset)
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.info = person.info | uint32(1<<gunOffset)
	}
}

func (p *GamePerson) HasGun() bool {
	flag := (p.info & uint32(1<<gunOffset)) >> gunOffset
	return flag == 1
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.info = person.info | uint32(1<<familyOffset)
	}
}

func (p *GamePerson) HasFamily() bool {
	flag := (p.info & uint32(1<<familyOffset)) >> familyOffset
	return flag == 1
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.info = person.info | uint32(personType)
	}
}

func (p *GamePerson) Type() int {
	mask := uint32(3)
	return int(p.info & mask)
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint16(respect)<<12
	}
}
func (p *GamePerson) Respect() int {
	return int(p.stats & uint16(statsBaseMask<<12) >> 12)
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint16(level)<<8
	}
}

func (p *GamePerson) Level() int {
	return int(p.stats & uint16(statsBaseMask<<8) >> 8)
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint16(strength)<<4
	}
}

func (p *GamePerson) Strength() int {
	return int(p.stats & uint16(statsBaseMask<<4) >> 4)
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint16(experience)
	}
}
func (p *GamePerson) Experience() int {
	return int(p.stats & uint16(statsBaseMask))
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.info = person.info | uint32(1<<houseOffset)
	}
}

func (p *GamePerson) HasHouse() bool {
	flag := (p.info & uint32(1<<houseOffset)) >> houseOffset
	return flag == 1
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
}

func (p *GamePerson) X() int {
	return int(p.xCoord)
}

func (p *GamePerson) Y() int {
	return int(p.yCoord)
}

func (p *GamePerson) Z() int {
	return int(p.zCoord)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}
