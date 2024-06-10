package datastructures

import "fmt"

// KeyValue represents a key-value pair
type KeyValue struct {
	key   string
	value string
}

// HashTable represents the hash table with chaining
type HashTable struct {
	buckets map[int][]KeyValue
}

// NewHashTable creates a new hash table
func NewHashTable() *HashTable {
	return &HashTable{buckets: make(map[int][]KeyValue)}
}

// HashFunction is a simple hash function
func HashFunction(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % 10 // Assuming we have 10 buckets
}

// Insert adds a key-value pair to the hash table
func (ht *HashTable) Insert(key string, value string) {
	index := HashFunction(key)
	ht.buckets[index] = append(ht.buckets[index], KeyValue{key, value})
}

// Search finds a value by its key in the hash table
func (ht *HashTable) Search(key string) (string, bool) {
	index := HashFunction(key)
	bucket := ht.buckets[index]
	for _, kv := range bucket {
		if kv.key == key {
			return kv.value, true
		}
	}
	return "", false
}

// PrintAll prints all key-value pairs in the hash table
func (ht *HashTable) PrintAll() {
	for index, bucket := range ht.buckets {
		fmt.Printf("Bucket %d:\n", index)
		for _, kv := range bucket {
			fmt.Printf("  Key: %s, Value: %s\n", kv.key, kv.value)
		}
	}
}

func MainHashTables() {
	hashTable := NewHashTable()

	// Insert toys into the hash table
	hashTable.Insert("car", "toy car")
	hashTable.Insert("robot", "toy robot")

	// Search for toys in the hash table
	value, found := hashTable.Search("car")
	if found {
		fmt.Println("Found:", value) // Found: toy car
	} else {
		fmt.Println("Car not found")
	}

	value, found = hashTable.Search("robot")
	if found {
		fmt.Println("Found:", value) // Found: toy robot
	} else {
		fmt.Println("Robot not found")
	}

	hashTable.PrintAll()
}
