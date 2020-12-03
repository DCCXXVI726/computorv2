NAME = computor

SRC_DIR = src
SRC = main.go complex.go createTokens.go 
COMP = go build

$(NAME): $(addprefix $(SRC_DIR)/, $(SRC)) 
	$(COMP) -o $(NAME) $(addprefix $(SRC_DIR)/, $(SRC)) 

all: $(NAME)

fclean:
	rm $(NAME)

re: fclean $(NAME)
