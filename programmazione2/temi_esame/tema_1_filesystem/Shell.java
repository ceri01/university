package filesystem;

import java.util.List;
import java.util.scanner;

/*	Questa classe rappresenta un shell, quindi un interprete che riceve dei comandi
 *	e li esegue. i comandi sono: (se tra quadre opzionale)
 *		ls [path]: che elenca il contenuto della directory indicato dal 
 *				   path o della directory corrente.
 *		size [path]: riporta la dimensione della entry indicata dal path
 *					 o della directory corrente
 *		mkdir path: crea una directory indicata dal path
 *		mkfile path size: crea il file indicato dal path, o nella radice del filesystem
 *		cd [path]: modifica la directory corrente in quella indicata dal path,o nella
 *				   radice del filesystem
 *		pwd: stampa il nome della directory corrente
 *
 *	Gli oggetti istanza di questa classe sono mutabili in quanto e' possibile cambiare la
 *	direcotry corrente in cui si trova la shell
 *
 *	ABS FUN = ABS() = 
 *
 *	REP INV = 
 *
 *
 */

public class Shell {
	// Attributi
	private Filesystem fs;
	private Path currentDirectory;

	// Costruttore
	private Shell(Filesystem fs) {
		if(fs == null) throw new IllegalArgumentException("Il filesystem ha valore nullo");
		this.fs = fs;
		this.currentDirectory = fs.root();
	}

	// Metodi

	public static void interpreter(String s) {
		switch()
	} 
		
		
	public static void main(String[] args) {
		Scanner s = new Scanner(System.in)	
		Filesystem fs = new Filesystem();
		for(;;) {
			
		}
	}
}

