package drivers


type LEDDriver interface {
    Write(colors interface{}) error
}
