package filesystem;

import java.util.Map;
import java.util.HashMap;

/*	La classe Directory rappresenta una cartella nel filesystem, essa e' mutabile in quanto il suo contenuto
 *	puo' cambiare (aggiungere file o cartelle)
 *
 *	ABS FUN: ABS(name, size, path, content) = per rappresentare la una directory utilizziamo una stringa per
 *											  indicare il nome (name), un valore intero per indicare la dimensione (size)
 *											  un attributo di tipo AbsolutePath che indica il percorso della entry (path)
 *											  e infine con una mappa di Entry indichiamo il contenuto della directory 
 *											  (content).
 *											  La mappa e' formata da chiave e valore, come chiave abbiamo la path del file o
 *											  directory, metre come valore avremo l'oggetto di tipo Entity (File o Directory)
 *
 *											  content = [pathFile1 -> file1
 *											  			 pathDirectory2 -> directory2
 *											  			 ...]
 *
 *	REP INV = name != null
 *			  size != null && size >= 0 && size = somma delle dimensioni del contenuto di content
 *			  path != null
 *			  content != null && content.get(path file o directory)!= null
 */

public class Directory implements Entry {
	
	// Attributi
	private final String name;
	private int size;
	private AbsolutePath path;
	private Map<String, Entry> content;

	//Costruttore
    public Directory(AbsolutePath path) {
        if(path == null) throw new IllegalArgumentException("Il percorso del file non puo' essere nullo");
		this.name = path.name();
        this.size = 0;
        this.path = path;
		this.content = new HashMap<String, Entry>();
    }}

	//Metodi
	public String name() {
		String name = new String(this.name);
		return name;
	}

	public int size() {
		int size = this.size;
		return size;
	}

	public Path path() {
		Path p = new AbsolutePath(this.path);
		return p;
	}

	/*	Metodo che permette l'inserimento di una Entry (Directory) nella mappa
	 *
	 * 	Pre-condizioni: path non deve essere nullo
	 *	Effetti collaterali: this.size cambia
	 *	Post-condizioni: viene inserita una directory nella mappa
	 *	
	 *	Questo metodo e' protected perche' verra' utilizzato dalla Filesystem, ma non
	 *	dall'utente.
	 */

	protected void insert(Path path) {
		
	}

	/*	Metodo che permette l'inserimento di una Entry (File) nella mappa
	 *
	 * 	Pre-condizioni: path non deve essere nullo, size >= 0 e non nullo
	 *	Effetti collaterali: this.size cambia
	 *	Post-condizioni: viene inserita un File nella mappa
	 *
	 *	Questo metodo e' protected perche' verra' utilizzato dal Filesystem, ma non
	 *	dall'utente.
	 */
	protected void insert(Path path, int size) {
	
	}


}
