package cache

// Создаю тип
// (наподобие справочник например контрагенты. у которго есть своя таблица)
type cache struct {
	table map[string]int
}

// привязываю функцию к типу (например как в модуле объекта)
// функция возвращает ссылку на обЪЕКТ а не само значение.
// по этому вначале создам контрагента а потом уже верну ссылку, при этом сама таблица
// должна быть хоть чем то заполнена За меня это делал конфиг. а здесь явно
// & - Это как ФУНКЦИЯ платформы ПолучитьСсылку()
// * - это как ФУНКЦИЯ ПолучитьОбъект()

func New() *cache {
	var c cache
	c.table = make(map[string]int)
	return &c
}

func (c *cache) Set(s string, val int) {
	c.table[s] = val
}

func (c *cache) Get(s string) int {
	return c.table[s]
}

func (c *cache) Delete(s string) {
	delete(c.table, s)
}
