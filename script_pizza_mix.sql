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

