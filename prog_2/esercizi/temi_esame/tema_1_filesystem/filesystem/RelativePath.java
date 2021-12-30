import java.util.List;
import java.util.ArrayList;
import java.util.Array;

/*	Questa classe estende Path e rappresenta un percorso relativo.
 *
 *	ABS FUN = ABS(path, parent) = Il percorso viene rappresentato tramite due attributi, path che indica la lista di nomi di cartelle per
 *								  per arrivare fino al file\cartella associato a questo path. E il parent rappresenta la cartella in cui
 *								  l'entita' relativa al percorso risiede.
 *
 *	REP INV = parent != null
 *			  path != null && path non contiene elementi nulli && il primo elemento di path non puo' essere "".
 *			  nessun elemento di path puo' contenere il carattere ""
 */

public class RelativePath extends Path {
	// Attributi
	
	private List<String> path;
	private String parent;
	
	// Costruttore
	
	public RelativePath(String path) {
		if(path == null) throw new IllegalArgumentException("Il percorso non puo' essere nullo");
		path = new ArrayList<String>();
		for(String s : path.split(":")) {
			if(s == "" && s == null) throw new IllegalArgumentException("Percorso non valido, una parte del percorso potrebbe essere nulla oppure una stringa non valida (stringa vuota)");
			this.path.add(s);
		}
		
		this.parent = this.path.get(this.size() - 2);
	}


	// Metodi

	public String name() {
		String s = new String(this.path.get(path.size() - 1));
		return s;
	}
	
	public String toString() {
		String s = new String(this.path.toString());
	}

	/*	Metodo che permette di recuperare il nome della cartella padre
	 *	
	 *	Post-condizioni: Viene restituita una stringa che indica il nome della cartella padre
	 */

	public String parent() {
		String s = new String(this.path.get(this.size() - 2));
		return s
	}
}
