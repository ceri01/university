package graph

/*
	Una seconda rappresentazione dei grafi è quella della matrice di adiacenza, ovvero si ha una matrice n*n in cui su righe
	sono presenti gli archi uscenti e sulle colonne gli archi entranti.
	Ogni elemento della matrice rappresenta un arco dal nodo indicato sulla riga al nodo indicato sulla colonna
	(questo per grafi orientati, mentre per i grafi non orientati la matrice sarà simmetrica).
	Il costo di questa rappresentazione in termini di spazio è nell'ordine di θ(n^2), perchè la matrice è n*n.
	La matrice d'incidenza è utile per fare rappresentazioni algebriche ma poco efficiente per quanto riguarda lo spazio.
*/
