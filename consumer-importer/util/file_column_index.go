package util

//HeaderSize size of the header of the file layout
const HeaderSize = 153

//CPFEnd latest index of CPF column in the file
const CPFEnd = 19

//PrivateEnd latest index of Private column in the file
const PrivateEnd = CPFEnd + 12

//IncompletoEnd latest index of Incompleto column in the file
const IncompletoEnd = PrivateEnd + 12

//DataUltimaCompraEnd latest index of 'Data ultima compra' column in the file
const DataUltimaCompraEnd = IncompletoEnd + 22

//TicketMedioEnd latest index of 'Ticket medio' column in the file
const TicketMedioEnd = DataUltimaCompraEnd + 22

//TicketUltimaCompraEnd latest index of 'Ticket ultima compra' column in the file
const TicketUltimaCompraEnd = TicketMedioEnd + 24

//LojaMaisFrequenteEnd latest index of 'Loja Mais Frequente' column in the file
const LojaMaisFrequenteEnd = TicketUltimaCompraEnd + 20
