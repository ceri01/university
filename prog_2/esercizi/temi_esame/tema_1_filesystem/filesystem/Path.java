import java.util.List;
import java.util.ArrayList;
import java.util.Array;

/*	Questa classe astratta rappresenta un generico percorso di un file o cartella. Estendendo la classe tramite 
 *	AbstractPath e RelativePath possiamo possiamo creare la distinzione tra percorso relativo e assoluto, che 
 *	differiscono nel fatto che il percorso assoluto parte dalla root, mentre quello relativo no
 */

public abstract class Path {
	private List<String> path;

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
		String s = new String(this.path.toString());
	}
}
