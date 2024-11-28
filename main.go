package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "fmt"
)

type Product struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    ImageURL    string `json:"image_url"`
    Price       int    `json:"price"`
    Brand       string `json:"brand"`
    Flavor      string `json:"flavor"`
    Ingredients string `json:"ingredients"`
}

type Order struct {
    ID         uint         `json:"id"`
    Products   []OrderItem  `json:"products"`
    TotalPrice int          `json:"total_price"`
    Date       string       `json:"date"`
}

type OrderItem struct {
    ProductID uint   `json:"product_id"`
    Name      string `json:"name"`
    Price     int    `json:"price"`
    Quantity  int    `json:"quantity"`
    ImageURL  string `json:"image_url"`
}

var orders = []Order{}
var nextOrderID uint = 1

var products = []Product{
    {ID: 1, Name: "Snickers", Description: "Шоколадный батончик с нугой, карамелью и орехами.", ImageURL: "https://media.istockphoto.com/id/529240903/ru/фото/snickers-шоколада-изолированные-на-белом-фоне.jpg?s=612x612&w=0&k=20&c=ySk6UIkmYIdvVo2M8FkZCGipu99YoGYP2DgTlVsaGFI=", Price: 50, Brand: "Mars", Flavor: "Шоколад с орехами", Ingredients: "Шоколад, карамель, арахис, нуга"},
    {ID: 2, Name: "M&M's", Description: "Шоколадные драже в цветной глазури.", ImageURL: "https://media.istockphoto.com/id/458135731/ru/фото/развернутый-m-ms-молочный-шоколад-конфеты.jpg?s=612x612&w=0&k=20&c=sQWEx_PiWZ5DoyhrMwzN7OcaeJ6vxC-3bf2DU8AiyGI=", Price: 70, Brand: "Mars", Flavor: "Шоколад", Ingredients: "Шоколад, сахар, красители"},
    {ID: 3, Name: "Toblerone", Description: "Швейцарский шоколад с медом и миндалем в виде треугольников.", ImageURL: "https://media.istockphoto.com/id/539951468/ru/фото/шоколадный-toblerone.jpg?s=612x612&w=0&k=20&c=nGyfE4Uq-X5XrWza4zRN3E_0Cf1ifKSfS8JRa45lzwU=", Price: 150, Brand: "Toblerone", Flavor: "Шоколад с миндалем и медом", Ingredients: "Шоколад, миндаль, мед"},
    {ID: 4, Name: "KitKat", Description: "Шоколадный батончик с вафлями внутри.", ImageURL: "https://media.istockphoto.com/id/458701489/ru/фото/kitkat-шоколадные-конфеты-бар-развернутый.jpg?s=612x612&w=0&k=20&c=ccwSWQVfy3kdg0tSOVNBTE7RRyT5CUl1ELA5dt0TReA=", Price: 40, Brand: "Nestlé", Flavor: "Шоколад с вафлями", Ingredients: "Шоколад, вафли, сахар, молоко"},
    {ID: 5, Name: "Ferrero Rocher", Description: "Шоколадные конфеты с орехом и начинкой из лесного ореха.", ImageURL: "https://media.istockphoto.com/id/460156521/ru/фото/ферреро-рошер-рождественская-ёлка.jpg?s=612x612&w=0&k=20&c=1n4DG6dWh-uHohKkNYEchFRjUaHvUVn6zKD20V3xeFI=", Price: 300, Brand: "Ferrero", Flavor: "Шоколад с орехом", Ingredients: "Шоколад, лесной орех, сахар, масло"},
    {ID: 6, Name: "Twix", Description: "Шоколадный батончик с карамелью и печеньем.", ImageURL: "https://media.istockphoto.com/id/458637643/ru/фото/twix-брусок-шоколада.jpg?s=612x612&w=0&k=20&c=3bg6beI4DZXtk7reRGo1e0e10ef4DKnOVJhXjTghvtU=", Price: 50, Brand: "Mars", Flavor: "Шоколад с карамелью и печеньем", Ingredients: "Шоколад, карамель, печенье, сахар"},
    {ID: 7, Name: "Milka", Description: "Шоколад с молочным вкусом и миндалем.", ImageURL: "https://media.istockphoto.com/id/526733402/ru/фото/милка-шоколадные-батончики.jpg?s=612x612&w=0&k=20&c=dGTWz_axK8-Ea9LnKOiUWlKZ-TVMqgs5U67Ku8w9oo4=", Price: 100, Brand: "Milka", Flavor: "Молочный шоколад с миндалем", Ingredients: "Шоколад, молоко, миндаль, сахар"},
    {ID: 8, Name: "Ritter Sport", Description: "Немецкий шоколад с цельным фундуком.", ImageURL: "https://media.istockphoto.com/id/507276506/ru/фото/вскоре-риттер-спорт-конфеты-мини.jpg?s=612x612&w=0&k=20&c=cB8SeSMDeOglwA2iR2GskPIslhOw2KefD73xjTondHU=", Price: 120, Brand: "Ritter Sport", Flavor: "Шоколад с фундуком", Ingredients: "Шоколад, фундук, сахар, молоко"},
    {ID: 9, Name: "Skittles", Description: "Жевательные конфеты с фруктовыми вкусами в яркой оболочке.", ImageURL: "https://media.istockphoto.com/id/507068144/ru/фото/кегли-candy.jpg?s=612x612&w=0&k=20&c=AecgihKETfh_GqgVX1R70vSVel1heO1_LslS-wqScdU=", Price: 60, Brand: "Mars", Flavor: "Фрукты", Ingredients: "Сахар, фруктовые ароматизаторы, красители"},
    {ID: 10, Name: "Lindt", Description: "Швейцарский темный шоколад с насыщенным вкусом.", ImageURL: "https://media.istockphoto.com/id/469433920/ru/фото/стек-линдор-конфет.jpg?s=612x612&w=0&k=20&c=lLV6M8_61SYkp6e6QKmS6aMLZY7IYTY6kq8iDqHRxqA=", Price: 250, Brand: "Lindt", Flavor: "Темный шоколад", Ingredients: "Какао, сахар, ваниль"},
}

var nextID uint = 11

func createOrder(c *gin.Context) {
    var incoming struct {
        Products   []OrderItem `json:"products"`
        TotalPrice int         `json:"total_price"`
    }

    if err := c.ShouldBindJSON(&incoming); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newOrder := Order{
        ID:         nextOrderID,
        Products:   incoming.Products,
        TotalPrice: incoming.TotalPrice,
        Date:       "2024-11-28", // Вы можете использовать текущую дату
    }
    nextOrderID++
    orders = append(orders, newOrder)

    c.JSON(http.StatusCreated, gin.H{"message": "Заказ сохранён", "order": newOrder})
}

func getAllOrders(c *gin.Context) {
    c.JSON(http.StatusOK, orders)
}

func getProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
        return
    }

    for _, product := range products {
        if product.ID == uint(id) {
            c.JSON(http.StatusOK, product)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func createProduct(c *gin.Context) {
    var product Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = nextID
    nextID++
    products = append(products, product)
    c.JSON(http.StatusCreated, product)
}

func updateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
        return
    }

    for i, product := range products {
        if product.ID == uint(id) {
            if err := c.ShouldBindJSON(&products[i]); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            products[i].ID = uint(id)
            c.JSON(http.StatusOK, products[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func deleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
        return
    }

    for i, product := range products {
        if product.ID == uint(id) {
            products = append(products[:i], products[i+1:]...)
            c.Status(http.StatusOK)
            fmt.Println("Product deleted:", product)
            return
        }
    }
    fmt.Println("Product not found with ID:", id)
    c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
}

func getAllProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    router := gin.Default()
    router.Use(CORSMiddleware())

    router.GET("/products", getAllProducts)           // Получить все продукты
    router.GET("/products/:id", getProductByID)       // Получить продукт по ID
    router.POST("/products", createProduct)           // Создать новый продукт
    router.PUT("/products/:id", updateProduct)        // Обновить продукт по ID
    router.DELETE("/products/:id", deleteProduct)     // Удалить продукт по ID
    router.POST("/orders", createOrder)               // Создать новый заказ
    router.GET("/orders", getAllOrders)

    router.Run(":8080")
}