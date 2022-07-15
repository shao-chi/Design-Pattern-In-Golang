package main

func main() {
	h_factory := hawaiianPizzaFactory{}
	m_factory := margheritaPizzaFactory{}

	hp_1 := h_factory.createProduct()
	hp_1.operate()

	mp_1 := m_factory.createProduct()
	mp_1.operate()

	hp_2 := h_factory.createProduct()
	hp_2.operate()

	mp_2 := m_factory.createProduct()
	mp_2.operate()

	hp_3 := h_factory.createProduct()
	hp_3.operate()

	mp_3 := m_factory.createProduct()
	mp_3.operate()
}
