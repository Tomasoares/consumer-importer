package util

const (
	//HeaderSize size of the header of the file layout
	HeaderSize = 157

	//CPFEnd latest index of CPF column in the file
	CPFEnd = 19

	//PrivateEnd latest index of Private column in the file
	PrivateEnd = CPFEnd + 12

	//IncompletoEnd latest index of Incompleto column in the file
	IncompletoEnd = PrivateEnd + 12

	//DataUltimaCompraEnd latest index of 'Data ultima compra' column in the file
	DataUltimaCompraEnd = IncompletoEnd + 22

	//TicketMedioEnd latest index of 'Ticket medio' column in the file
	TicketMedioEnd = DataUltimaCompraEnd + 22

	//TicketUltimaCompraEnd latest index of 'Ticket ultima compra' column in the file
	TicketUltimaCompraEnd = TicketMedioEnd + 24

	//LojaMaisFrequenteEnd latest index of 'Loja Mais Frequente' column in the file
	LojaMaisFrequenteEnd = TicketUltimaCompraEnd + 20
)
