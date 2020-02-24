package structures

// структура сообщения, которое мы будем отправлять на сервер
type Message struct {
	Identifier int    // идентификатор сообщения
	Message    string // команда, которую мы отправим по сокетам на сервер
	Name       string // тип протокола. по умолчанию WebRcon
}

// структура сообщения, которое приходит с сервера
type Response struct {
	Identifier int // мне нужен только идентификатор сообщения
}

// структура конфига
type Configuration struct {
	Host     string // адрес сервера вида IP:Port
	Password string // RCON-пароль
}
