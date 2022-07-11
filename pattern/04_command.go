package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Паттерн Command относится к поведенческим паттернам уровня объекта.
// Паттерн Command позволяет представить запрос в виде объекта.
// Из этого следует, что команда - это объект.
// Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
// Единственное, что должен знать инициатор, это как отправить команду.

//+:	Убирает прямую зависимость между объектами, вызывающими операции, и объектами,
//которые их непосредственно выполняют.
//	Позволяет реализовать простую отмену и повтор операций.
//	Позволяет реализовать отложенный запуск операций.
//	Реализует принцип открытости/закрытости.
//-:	Усложняет код программы из-за введения множества дополнительных классов.

// Command представляет собой интерфейс
type Command interface {
	Execute() string
}

// ToggleOnCommand и ToggleOffCommand содержат в себе запросы к Receiver, которые тот должен выполнять.
// В свою очередь Receiver содержит только набор действий,
// которые выполняются при обращении к ним из ToggleOnCommand и ToggleOffCommand .

// ToggleOnCommand реализует интерфейс Command
type ToggleOnCommand struct {
	receiver *Receiver
}

// Execute какая-то команда.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реализует интерфейс Command
type ToggleOffCommand struct {
	receiver *Receiver
}

// Execute какая-то команда.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Receiver - получатель команды
type Receiver struct {
}

// ToggleOn объект, реализующий команду
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// ToggleOff  объект, реализующий команду
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

//Invoker умеет складывать команды в стопку и инициировать их выполнение по какому-то событию.
//Обратившись к Invoker можно отменить команду, пока та не выполнена.

// Invoker - отправитель команды
type Invoker struct {
	commands []Command
}

// StoreCommand - добавляет команды в очередь
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand удаляет команды из очереди
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute выполняет все команды в очереди
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}
