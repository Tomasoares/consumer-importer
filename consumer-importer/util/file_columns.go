package util

const HEADER_SIZE = 153;

const CPF_END = 19;
const PRIVATE_END = CPF_END + 12;
const INCOMPLETO_END = PRIVATE_END + 12;
const DATA_ULTIMA_COMPRA_END = INCOMPLETO_END + 22;
const TICKET_MEDIO_END = DATA_ULTIMA_COMPRA_END + 22;
const TICKET_ULTIMA_COMPRA_END = TICKET_MEDIO_END +24;
const LOJA_MAIS_FREQUENTE_END = TICKET_ULTIMA_COMPRA_END + 20;