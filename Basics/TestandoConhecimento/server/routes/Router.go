package routes

import (
	"example.com/teste/server/routes/handlers"
	"github.com/gin-gonic/gin"
)

func PdvRoutes(server *gin.Engine, handler *handlers.PdvHandlers) {
	server.POST("/PDV/pesquisarPdvPorId", handler.PesquisarPdvPorId)
	server.POST("/PDV/pesquisarPdvPorNome")
	server.POST("/PDV/pegarIdPDV")
	server.POST("/PDV/inserirPDV")
	server.POST("/PDV/editarPDV")
	server.POST("/PDV/deletarPDV")
}

func EntradaRoutes(server *gin.Engine) {
	server.POST("/ENTRADA/pegarIdPEntrada")
	server.POST("/ENTRADA/pesquisarNotaLancar")
	server.POST("/ENTRADA/pesquisarFornecedorPEntrada")
	server.POST("/ENTRADA/pesquisarProdutoEntPorId")
	server.POST("/ENTRADA/inserirEntradas")
	server.POST("/ENTRADA/pesquisarEntradaPorId")
	server.POST("/ENTRADA/editarEntrada")
	server.POST("/ENTRADA/deletarEntrada")
	server.POST("/ENTRADA/salvaDados")
}

func FornecedorRoutes(server *gin.Engine) {
	server.POST("/FORNECEDOR/pesquisarFornecedorPorId")
	server.POST("/FORNECEDOR/pesquisarFornecedorPorNome")
	server.POST("/FORNECEDOR/editarFornecedor")
	server.POST("/FORNECEDOR/pegarIdpFornecedor")
	server.POST("/FORNECEDOR/inserirFornecedor")
	server.POST("/FORNECEDOR/deletarFornecedor")
	server.POST("/FORNECEDOR/cepFornecedor")

}

func ProdutoRoutes(server *gin.Engine) {
	server.POST("/PRODUTO/pesquisarProdutoPorId")
	server.POST("/PRODUTO/pesquisarProdutoPorNome")
	server.POST("/PRODUTO/pegarIdProduto")
	server.POST("/PRODUTO/inserirProduto")
	server.POST("/PRODUTO/editarProduto")
	server.POST("/PRODUTO/deletarProduto")
	server.POST("/PRODUTO/verificarAutomacaoProduto")
	server.POST("/PRODUTO/verificarOrdem")
}

func VendedorRoutes(server *gin.Engine) {
	server.POST("/VENDEDOR/pesquisarPorNome")
	server.POST("/VENDEDOR/pesquisarPorId")
	server.POST("/VENDEDOR/editarVendedor")
	server.POST("/VENDEDOR/pegarImagemAntiga")
	server.POST("/VENDEDOR/inserirDados")
	server.POST("/VENDEDOR/cep")
	server.POST("/VENDEDOR/deletar")
	server.POST("/VENDEDOR/pegarIdpChave")
	server.POST("/VENDEDOR/recuperarAcesso")
	server.POST("/VENDEDOR/VerificarCpf")
}

func DashboardRoutes(server *gin.Engine) {
	server.POST("/DASHBOARD/pegarVenOperacao")
	server.POST("/DASHBOARD/pegarOperacaoAtivo")
	server.POST("/DASHBOARD/valorMedioTickets")
	server.POST("/DASHBOARD/valorMaximoTickets")
	server.POST("/DASHBOARD/valorMinimoTickets")
}

func MovimentoRoutes(server *gin.Engine) {
	server.POST("/MOVIMENTO/pesquisarVendedorMovimento")
	server.POST("/MOVIMENTO/pegarProdutosMovimento")
	server.POST("/MOVIMENTO/pegarUltimoIdMovimento")
	server.POST("/MOVIMENTO/pesquisarPorNomeMovimento")
	server.POST("/MOVIMENTO/pesquisarTicketsMovimento")
	server.POST("/MOVIMENTO/pesquisarTicketsMovimentoEncerrar")
	server.POST("/MOVIMENTO/registrarTicket")
	server.POST("/MOVIMENTO/registrarDetalhesTicket")
	server.POST("/MOVIMENTO/fecharTickets")
	server.POST("/MOVIMENTO/pacoteCheioMovimento")
}

func SorteioRoutes(server *gin.Engine) {
	server.POST("/SORTEIO/nomeVendedorSorteio")
	server.POST("/SORTEIO/preencherPdv")
	server.POST("/SORTEIO/salvarRecuperarSenha")
	server.POST("/SORTEIO/pesquisarIdVendedorSorteio")
	server.POST("/SORTEIO/cadastrarSorteio")
	server.POST("/SORTEIO/preencherAtribuicao")
	server.POST("/SORTEIO/pegarIdSorteio")
	server.POST("/SORTEIO/pesquisarSorteio")
	server.POST("/SORTEIO/nomesSorteio")
	server.POST("/SORTEIO/deletarSorteio")
}
