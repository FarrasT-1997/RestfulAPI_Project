FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./

COPY . .

# COPY ./config/config.go ./

# COPY ./constant/constant.go ./

# COPY ./controller/cartController.go ./
# COPY ./controller/paymentMethod.go ./
# COPY ./controller/product&category.go ./
# COPY ./controller/transaction.go ./
# COPY ./controller/userController.go ./

# COPY ./lib/database/cart.go ./
# COPY ./lib/database/payment.go ./
# COPY ./lib/database/product.go ./
# COPY ./lib/database/transactions.go ./
# COPY ./lib/database/user.go ./

# COPY ./middlewares/jwt.go ./
# COPY ./middlewares/logmiddlewares.go ./

# COPY ./models/category.go ./
# COPY ./models/payment-method.go ./
# COPY ./models/product-description.go ./
# COPY ./models/product.go ./
# COPY ./models/shopping-cart.go ./
# COPY ./models/transaction.go ./
# COPY ./models/user.go ./

# COPY ./routes/routes.go ./

RUN go build -o program

CMD ./program