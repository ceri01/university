import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.StringBuilder;
import java.util.Iterator;

/*	Questa classe astratta rappresenta un generico percorso di un file o cartella.
 *	Il percorso puo' essere relativo o assoluto, in piu' il percorso e' immutabile
 *	in quanto non e' previsto che una Entry cambi posizione, quindi il suo path non
 *	cambia.
 *	
 *	ABS FUN = ABS(path, isAbsolute) = Il percorso viene rappresentato tramite una lista di 
 *									  stringe (path) e da isAbsolute che indica se il percorso e'
 *									  assoluto o no (true -> assoluto, false -> relativo).
 *
 *	REP INV = path != null
 *			  ogni stringa in path !=  null
 *
 */

public class Path {
	// Attributi
	
	private List<String> path;
	private boolean isAbsolute;
	
	// Costruttore
	
	private Path(final List<String> path, final boolean isAbsolute) {
		if(path.isEmpty()) throw new IllegalArgumentException("path e' vuoto e non puo' rappresentare un percorso");
		for(String s : path) {
			if(s.isEmpty()) throw new IllegalArgumentException("Non e' possibile avere una stringa vuota nel percorso");
		}
		this.path = List.copyOf(path);
		this.isAbsolute = isAbsolute;
	}

	// Metodi

	/*	Metodo che permette di ricavare il nome del file o della cartella
	 *
	 *	Post-condizioni: Ritorna una stringa che rappresenta il nome del file o della cartella
	 */

	public String name() {
		String s = new String(this.path.get(path.size() - 1));
		return s;
	}

	/*	Metodo che permette di ricavare il percorso del file o della cartella
	 *
	 *	Post-condizioni: Ritorna una stringa che rappresenta il percorso del file o della cartella
	 */

	public String toString() {
		StringBuilder str = new StringBuilder("")
		if(this.isAbsolute){
			str.append(":");
		}
		Iterator i = this.path.iterator();
		for(i.hasNext()) {
			str.append(i.Next());
			if(i.hasNext()) {
				str.append(":");
			}
		}
		return str.toString()
	}
}
