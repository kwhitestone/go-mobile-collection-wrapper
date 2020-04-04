package syncMapWrapper

const keyObjectValuePointerTemplate = `{{range .Types}}
type {{.KeyTitle}}{{.Value}}Map struct {
	sync.Mutex
	s map[{{.Key}}]*{{.Value}}
}

func New{{.KeyTitle}}{{.Value}}Map() *{{.KeyTitle}}{{.Value}}Map {
	return &{{.KeyTitle}}{{.Value}}Map{}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Clear() {
	m.Lock()	
	m.s = make(map[{{.Key}}]*{{.Value}})
	m.Unlock()
}

func (m *{{.KeyTitle}}{{.Value}}Map) Equal(rhs *{{.KeyTitle}}{{.Value}}Map) bool {
	m.Lock()
	defer m.Unlock()
	if rhs == nil {
		return false
	}

	if len(m.s) != len(rhs.s) {
		return false
	}

	for k := range m.s {
		if m.s[k] != rhs.s[k] {
			return false
		}
	}

	return true
}

func (m *{{.KeyTitle}}{{.Value}}Map) MarshalJSON() ([]byte, error) {
	m.Lock()
	defer m.Unlock()
	return json.Marshal(m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) UnmarshalJSON(data []byte) error {
	m.Lock()
	defer m.Unlock()
	return json.Unmarshal(data, &m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) Copy(rhs *{{.KeyTitle}}{{.Value}}Map) {
	m.Lock()
	defer m.Unlock()
	m.s = make(map[{{.Key}}]*{{.Value}})
	for k, v := range rhs.s {
		m.s[k] = v
	}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Clone() *{{.KeyTitle}}{{.Value}}Map {
	m.Lock()
	defer m.Unlock()
	nS := make(map[{{.Key}}]*{{.Value}})
	for k, v := range m.s {
		nS[k] = v
	}
	return &{{.KeyTitle}}{{.Value}}Map{
		Mutex: sync.Mutex{},
		s: nS,
	}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Key(rhs *{{.Value}}) {{.Key}} {
	m.Lock()
	defer m.Unlock()
	for i, lhs := range m.s {
		if lhs == rhs {
			return i
		}
	}
	return reflect.Zero(reflect.TypeOf(m.s).Key()).Interface().({{.Key}})
}

func (m *{{.KeyTitle}}{{.Value}}Map) Set(key {{.Key}}, n *{{.Value}}) {
	m.Lock()
	defer m.Unlock()
	m.s[key] = n
}

func (m *{{.KeyTitle}}{{.Value}}Map) Remove(key {{.Key}}) {
	m.Lock()
	defer m.Unlock()
	delete(m.s, key)
}

func (m *{{.KeyTitle}}{{.Value}}Map) Count() int {
	m.Lock()
	defer m.Unlock()
	return len(m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) At(key {{.Key}}) *{{.Value}} {
	m.Lock()
	defer m.Unlock()
	u, ok := m.s[key]
	if !ok {
		return nil
	}
	return u
}

func (m *{{.KeyTitle}}{{.Value}}Map) Keys() *{{.KeyTitle}}Slice {
	m.Lock()
	defer m.Unlock()
	keys := New{{.KeyTitle}}Slice()
	for k := range m.s {
		keys.Append(k)
	}
	return keys
}

func (m *{{.KeyTitle}}{{.Value}}Map) objectsMap() map[{{.Key}}]{{.Value}} {
	m.Lock()
	defer m.Unlock()
	res := make(map[{{.Key}}]{{.Value}})
	for k, v:= range m.s {
		if v == nil {
			continue
		}
		res[k] = *v
	}
	return res
}

func new{{.KeyTitle}}{{.Value}}MapWithObjects(objects map[{{.Key}}]{{.Value}}) *{{.KeyTitle}}{{.Value}}Map {	
	pMap := make(map[{{.Key}}]*{{.Value}})
	for k, v := range objects {
		v1 := v
		pMap[k] = &v1
	}
	return &{{.KeyTitle}}{{.Value}}Map{sync.Mutex{}, pMap}
}

{{end}}`
