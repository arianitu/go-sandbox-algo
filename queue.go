package algo

type Element struct {
	data string
}

type Queue struct {
	elements []*Element
}

func NewQueue() *Queue {
	return &Queue {  elements: make([]*Element, 0)}
}

func (q *Queue) Push(data string) {
	q.elements = append(q.elements, &Element{data: data})
}

func (q *Queue) Pop() *Element {
	if q.IsEmpty() {
		return nil
	}
	
	returnValue := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	
	return returnValue
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	return len(q.elements)
}
