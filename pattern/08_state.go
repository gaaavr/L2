package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Паттерн State относится к поведенческим паттернам уровня объекта.
// Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и
// является объектно-ориентированной реализацией конечного автомата. Поведение объекта изменяется настолько,
// что создается впечатление, будто изменился класс объекта.
// Паттерн должен применяться:
// 1.Когда поведение объекта зависит от его состояния.
// 2.Поведение объекта должно изменяться во время выполнения программы.
// 3.Состояний достаточно много и использовать для этого условные операторы,
// разбросанные по коду, достаточно затруднительно.

// Плюсы:
//	1.Избавляет от множества условных операторов
//	2.Упрощает код контекста
//	Минусы:
//	1.Может усложнить код если состояний много и редко меняются

// MobileAlertStater представляет собой общий интерфейс для различных состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует громкость звонка в зависимости от состояния
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку с содержанием громкости звонка
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState меняет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert функция-конструктор для MobileAlert.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует звонок вибрацией
type MobileAlertVibration struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong реализует звонок звуком
type MobileAlertSong struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}
