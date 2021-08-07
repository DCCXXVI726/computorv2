NAME = computor

SRC_DIR = src
SRC = operation.go 
COMP = go build

$(NAME): $(addprefix $(SRC_DIR)/, $(SRC)) 
	$(COMP) -o $(NAME) $(addprefix $(SRC_DIR)/, $(SRC)) 

test: 
	go test $(addprefix ./,$(SRC_DIR))

all: $(NAME)

fclean:
	rm $(NAME)

re: fclean $(NAME)
