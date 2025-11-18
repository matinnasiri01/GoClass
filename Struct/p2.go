package main

import "fmt"

type CPU struct {
    Brand string
    Core  int
    Speed float32
}
type Memory struct {
    Capacity       int
    FrequencyInMHz int
}
type Storage struct {
    Capacity int
    Type     string
}
type Computer struct {
    cpu     CPU
    memory  Memory
    storage Storage
}


/// Composition
func struct2() {
    c := CPU{Brand: "Intel", Core: 8, Speed: 4.5}
    m := Memory{Capacity: 16, FrequencyInMHz: 3200}
    s := Storage{Capacity: 500, Type: "SSD"}

    comp := Computer{cpu: c, memory: m, storage: s}

    fmt.Printf("%+v\n", comp)
    fmt.Println("Core : ", comp.cpu.Core)
    fmt.Println("Memory Capacity : ", comp.memory.Capacity)
}

/// Embedding
type Product struct {
    Name  string
    Price int
}
type Electrical struct {
    Product
    WarrantyMonths int
}
type Clothing struct {
    Product
    Size     string
    Material string
}
func embedding() {
    cl := Clothing{}
    cl.Name = "Shirt"
    cl.Price = 300
    cl.Size = "S"
    cl.Material = "Cotton"
    fmt.Printf("%+v\n", cl)

    el := Electrical{Product: Product{Name: "Lamp", Price: 40}, WarrantyMonths: 12}
    fmt.Printf("%+v\n", el)
}