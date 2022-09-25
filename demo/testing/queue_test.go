package queue

// testing functionality of queue file
// to test use g test -v ./demo/testing

import "testing"

// t *testing.T is required for it to be reckognised as a test function
func TestCreateQueue(t *testing.T) {
	q := New(3)

	// for loop to add items to the queue
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect queue length: %v, expected %v", len(q.items), i)
		}

		// q.Append(i) will run and is expected to return true if run successfully
		if !q.Append(i) {
			t.Errorf("failed to append to queue: %v", i)
		}
	}

	// queue should be full at this point so we'll try to add something and see if we can
	if q.Append(4) {
		t.Errorf("should not be able to append to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	//retrieve items out of the queue
	for i := 0; i < 3; i++ {
		//Next() returns a bool and an int
		item, ok := q.Next()

		if !ok {
			t.Error("should be able to retrieve items from the queue")
		}

		if item != i {
			t.Errorf("not retrieving items in a FIFO method, got %v, expected %v", item, i)
		}
	}

	// all items should be retrieved out of the queue at this stage
	item, ok := q.Next()

	if ok {
		t.Errorf("should not be any more items in the queue, got %v", item)
	}
}
