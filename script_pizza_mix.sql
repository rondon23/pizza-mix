CREATE DATABASE pizza_mix;
USE pizza_mix;

DROP TABLE IF EXISTS `produto`;
CREATE TABLE `produto` (
  `produto_id` int(11) NOT NULL AUTO_INCREMENT,
  `nome_produto` varchar(255) NOT NULL,
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
   `sub_total` decimal(15,2) not null,
   primary key (`carrinho_id`),
   key `produto_FK` (`produto_id`),
   constraint `produto_FK` foreign key (`produto_id`) references `produto` (`produto_id`)
) ENGINE=InnoDB auto_increment=2006 DEFAULT CHARSET= latin1;
INSERT INTO `carrinho` VALUES
	(2000, 2000, 30.3, 5, 151.5, 'Pizza Frango com Catupiry', 151.5),
	(2001, 2001, 31.5, 5, 151.5, 'Portuguesa', 151.5),
	(2002, 2002, 31.5, 5, 157.5, 'Calabresa', 157.5);
	
drop table if exists `funcionario`;
create table `funcionario`(
	`funcionario_id` int(11) not null auto_increment,
	`nome` varchar(255) not null,
	`rg` varchar(14) not null,
	`cpf` varchar(14) not null,
	`rua` varchar(100) not null,
	`numero` int(11) not null,
	`bairro` varchar(100) not null,
	`cargo` varchar(200) not null,
	primary key (`funcionario_id`)
)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
insert  into `funcionario` values 
	(1, 'Fulando da Silva Sauro', '009648634', '03876378252', 'Rua das Galeras', 67, 'jardim macaubas', 'Padeiro');
	
drop table if exists `cliente`;
create table `cliente`(
	`cliente_id` int(11) not null auto_increment,
	`nome` varchar(255) not null,
	`cpf` varchar(14) not null,
	`endereco` varchar(255) not null,
	`telefone` varchar(14) not null,
	`numero` int(11) not null,
	`bairro` varchar(100) not null,
	`email` varchar(100) not null,
	primary key (`cliente_id`)
)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
insert into `cliente` values 
	(1, 'Ciclano Buenno', '03797478332', 'Rua Tiberio', '67998747834', 78, 'Jardim das Primaveras', 'ciclano@email.com');

drop table if exists `venda`;
create table `venda`(
	`venda_id` int(11) not null auto_increment,
	`nome` varchar(255) not null,
	`cpf` varchar(14) not null,
	`endereco` varchar(255) not null,
	`telefone` varchar(14) not null,
	`numero` int(11) not null,
	`bairro` varchar(255) not null,
	`email` varchar(100) not null,
	primary key (`venda_id`)
)ENGINE=InnoDB auto_increment= 2006 default CHARSET=latin1;
insert into `venda` values 
	(1, 'Fulano Silva', '03764589443', 'Rua Pirapozinho', '67996478334', 67, 'Jardim Marajoara', 'teste@teste.com');
	
drop table if exists `fornecedor`;
create table `fornecedor`(
	`fornecedor_id` int(11) not null auto_increment,
	`nome` varchar(255) not null,
	`endereco` varchar(255) not null,
	`numero` int(11) not null,
	`bairro` varchar(255) not null,
	`cnpj` varchar(18) not null,
	primary key (`fornecedor_id`)
	)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `fornecedor` values 
		(1, 'Paes e CIA Ltda', 'Rua Pioneiros', 987, 'Jardim Autonomista', '84839483749383');

drop table if exists `compra`;
create table `compra`(
	`compra_id` int(11) not null auto_increment,
	`data_compra` date not null,
	`fornecedor_id` int(11) not null,
	`produto_id` int(11) not null,
	`preco_unitario` decimal(15, 2) not null,
	`valor_compra` decimal(15, 2) not null,
	primary key (`compra_id`),
	key `produto_compra_FK` (`produto_id`),
    constraint `produto_compra_FK` foreign key (`produto_id`) references `produto` (`produto_id`),
    key `fornecedor_compra_FK` (`fornecedor_id`),
    constraint `fornecedor_compra_FK` foreign key (`fornecedor_id`) references `fornecedor` (`fornecedor_id`)
	)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `compra` values 
		(1, '2022-06-20 00:20:23', 1, 2000, 19.3, 20.6);
		
drop table if exists `contas_pagar`;
create table `contas_pagar`(
		`conta_pagar_id` int(11) not null auto_increment,
		`date` date not null,
		`descricao` varchar(255) not null,
		`valor` decimal(15, 2) not null,
		`compra_id` int(11) not null,
		primary key (`conta_pagar_id`),
		key `compra_FK` (`compra_id`),
		constraint `compra_FK` foreign key (`compra_id`) references `compra` (`compra_id`)
	)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `contas_pagar` values
		(1, '2022-06-20 00:20:23', 'Fermentos', 19.23, 1);
		
drop table if exists `contas_a_pagar`;
create table `contas_a_pagar`(
		`conta_a_pagar_id` int(11) not null auto_increment,
		`data` date not null,
		`valor` decimal(15, 2) not null,
		primary key (`conta_a_pagar_id`)
		)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `contas_a_pagar` values 
	(1, '2022-06-20 00:20:23', 19.2);
	
drop table if exists `mov_conta`;
create table `mov_conta` (
	`mov_conta_id` int(11) not null auto_increment,
	`valor` decimal(15, 2) not null,
	primary key (`mov_conta_id`)
)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `mov_conta` values 
	(1, 26.6);
	
drop table if exists `conta`;
create table `conta`(
	`conta_id` int(11) not null auto_increment,
	`data_movimentacao` date not null,
	`saldo` decimal(15, 2) not null,
	`numero_conta` varchar(14) not null,
	primary key (`conta_id`)
	)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `conta` values 
	(1, '2022-06-20 00:20:23', 17.8, '1459845');
	
drop table if exists `banco`;
create table `banco` (
	`banco_id` int(11) not null auto_increment,
	`endereco` varchar(255) not null,
	`numero` int(11) not null,
	`bairro` varchar(255) not null,
	`telefone` varchar(14) not null,
	primary key (`banco_id`)
	)ENGINE=InnoDB auto_increment=2006 default CHARSET=latin1;
	insert into `banco` values 
	(1, 'Rua Palmeiras', 14, 'Jardim Palmares', '6789498849');