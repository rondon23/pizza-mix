CREATE DATABASE pizza_mix;
USE pizza_mix;

DROP TABLE IF EXISTS `produto`;
CREATE TABLE `produto` (
  `produto_id` int(11) NOT NULL AUTO_INCREMENT,
  `nome_produto` varchar(100) NOT NULL,
  `descricao` varchar(100) NOT NULL,
  `valor_produto` decimal(15,2) not null,
  `foto_produto` blob not null,
  PRIMARY KEY (`produto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
INSERT INTO `produto` VALUES
	(2000,'Pizza Frango com Catupiry','Contem um delicioso frango com catupiri', 30.3 ,'teste'),
	(2001,'Portuguesa','Contem uma deliciosa pizza portuguesa', 31.5 ,'teste'),
	(2002,'Calabresa','Contem uma deliciosa pizza de calabresa', 31.5 ,'teste');

drop table if exists `carrinho`;
create table `carrinho`(
   `carrinho_id` int(11) not null auto_increment,
   `produto_id` int(11) not null,
   `preco` decimal(15,2) not null,
   `quantidade` int(11) not null,
   `total` decimal(15,2) not null,
   `nome_produto` varchar(100) not null,
   `subtotal` decimal(15,2) not null,
   primary key (`carrinho_id`),
   key `produto_FK` (`produto_id`),
   constraint `produto_FK` foreign key (`produto_id`) references `produto` (`produto_id`)
) ENGINE=InnoDB auto_increment=2006 DEFAULT CHARSET= latin1;
INSERT INTO `carrinho` VALUES
	(2000, 2000, 30.3, 5, 151.5, 'Pizza Frango com Catupiry', 151.5),
	(2001, 2001, 31.5, 5, 151.5, 'Portuguesa', 151.5),
	(2002, 2002, 31.5, 5, 157.5, 'Calabresa', 157.5);