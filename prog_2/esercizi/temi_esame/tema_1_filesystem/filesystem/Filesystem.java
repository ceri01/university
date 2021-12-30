package filesystem;

/*	La classe Filesystem rappresenta la parte principale dell'esercizio, ovvero il filesystem stesso.
 *	
 *	Un filesystem viene rappresentato tramite la root, ovvero la Directoy principale del filesystem,
 *	quindi per questo motivo un istanza della classe filesystem e' immutabile in quanto abbiamo solo la 
 *	directory root e non puo' essere cambiata.
 *
 *	Un istanza della classe Filesystem e' in grado di reperire un entry dato il suo path assoluto (getEntry), creare directory
 *	dato il path assoluto (createDirectory) creare un file dato il path assoluto e la dimensione (createFile), elencare il contenuto
 *	di una directory dato il suo path assoluto (listContent) e infine di ottenere la dimensione di una entry dato il path assoluto (getSize)
 *	
 *
 *	ABS FUN: ABS(root) = possiamo vedere il filesystem come un albero orientato che ha come radice la root. possiamo 
 *						 vedere i path come se fossero gli archi che collegano i vari nodi, e i nodi sono rappresentati dai
 *						 file e dalle Directory.
 *
 *	REP INV: root != NULL
 */


public class Filesystem {
	
	// Attributi
	private final Directory root;

	//Costruttore
	public Filesystem() {
		this.root = new Directory() ;
	}

	//Metodi
	
	/*	Metodo che permette di reperire una file o una directory dato il suo path assoluto
	 *
	 *	Pre-condizioni: path != null
	 *	Post-condizioni: viene restituira una entry (un file o una directory)
	 */

	public Entry getEntry(AbsolutePath path) {
		if(path == null) throw new IllegalArgumentException("Path nullo, impossibile reperire alcuna entry.");
	}
	
	/*	Metodo che permette di creare una directory
	 *
	 *	Pre-condizioni: path != null
	 *					path non deve essere gia' esistente.
	 *	Effetti collaterali: dimensione directory padre aumentata
	 *	Post-condizioni: Viene creata una directory nel path specificato
	 */
	
	public void createDirectory(AbsolutePath path) {
		
	}

	
	/*	Metodo che permette di creare un file
	 *
	 *	Pre-condizioni: path != null
	 *					path non deve essere gia' esistente.
	 *	Effetti collaterali: dimensione directory padre aumentata
	 *	Post-condizioni: Viene creato un file nel path specificato
	 */
	
	public void createFile(AbsolutePath path, int dim) {
		
	}

	/*	Metodo che permette di elencare il contenuto di una directory
	 *
	 *	Pre-condizioni: path != null
	 *					path deve specificare una directory e non un file
	 *	Post-condizioni: viene stampato il contenuto della directory
	 */
	
	public void listContent (AbsolutePath path) {
		
	}

	/*	Metodo che permette di ricavare la dimensione di una entry 
	 *	
	 *	Pre-condizioni: path != null
	 *	Post-condizioni: viene restituita la dimensione della entry specificata da path 
	 */
	
	public int getSize(AbsolutePath path) {
		
	}
}
