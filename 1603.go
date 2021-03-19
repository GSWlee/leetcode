type ParkingSystem struct {
    Big int
    Mid int
    Small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big,medium,small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType{
        case 1:
            if this.Big>0{
                this.Big--
                return true
            }
            return false
        case 2:
            if this.Mid>0{
                this.Mid--
                return true
            }
            return false
        case 3:
            if this.Small>0{
                this.Small--
                return true
            }
            return false
    }
    return false
}