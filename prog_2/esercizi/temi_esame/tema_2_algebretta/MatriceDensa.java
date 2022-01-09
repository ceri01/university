package tema_2_algebretta

import java.util.List;
import java.util.ArrayList;

/*	Questa classe rappresenta una generica matrice densa, la matrice non e' mutabile in quanto non
 *	e' richiesta la rimozione o aggiunta di elemetni nella matrice.
 *
 *	Per semplicita le matrici saranno matrici quadrate, come richiesto nella consegna.
 *
 *	ABS FUN = ABS(mat, dim) = mat e' un array che verra riempito di vettori, quindi mat rappresentera'
 *							  la matrice, e dim rappresenta la dimensione della matrice (numero righe e 
 *							  colonne della matrice).
 *							  (essendo matrici quadrate ci basta un intero per rappresentare la dimensione).
 *
 *
 *	REP INV = mat != null
 *			  dim > 0
 *	
 *	Non sara' possibile creare una matrice  formata da vettori vuoti in qunato la dimensione
 *	della matrice dovrebbe essere 0, 0 (quindi dim = 0) ma se solo mettiamo un vettore vuoto la dimensione della
 *	matrice sarebbe 0, 1, e non rispetterebbe la condizione che dice che le matrici devono essere quadrate.
 */


public class MatriceDensa implements Matrice {
	// Attributi
	private final Vettore[] mat;
	private final int dim; // se non ci fosse il vincolo delle matrici quadrate avrei usato un arrey di due interi.

	// Costruttore
	
	/*	Costruttore di MatriceDensa
	 *
	 *	pre-condizioni: l != null
	 *					l non contiene valori nulli
	 *					l non contiene vettori vuoti
	 *					l deve contenere vettori di dimensione uguale al numero di vettori in l
	 *	effetti collaterali: this.mat e this.dim possono venir modificati
	 *	post-condizioni: viene creato l'oggetto istanza di MatriceDensa
	 */
	
	public MatriceDensa(List<Vettore> l) {
		if(l == null) throw new IllegalArgumentException("La matrice non puo' essere nulla.")
		for(Vettore v : l) {
			if(v == null || v.length == 0) throw new IllegalArgumentException("Non e' possibile inserire un vettore nullo o vuoto, la consizione per cui la matrice deve essere quadrata non verrebbe rispettata.");
			if(l.length() != v.dim()) throw new IllegalArgumentException("il Vettore " + v.toString() + " non rispetta le condizioni per cui la matrice deve essere quadrata.");
		}
		this.mat = new int[l.length()];
		l.toArray(this.mat);
		this.dim = this.mat.length;
	}
	
	/*	Costruttore copia di MatriceDensa
	 *
	 *	per-condizioni: other != null
	 *	post-condizioni: crea un oggetto che e' la copia di other
	 */

	public MatriceDensa(MatriceDensa other) {
		if(other == null) throw new IllegalArgumentException("Impossibile creare una copia di una matrice che ha valore null.");
		for(int i = 0; i < m.dim(); i++) {
			Vettore tmp = new VettoreDenso(other.vet[i]);
			this.mat[i] = tmp;
		}
		this.dim = tmp.dim;
	}
	// Metodi
	
    public int[] dim() {
		int dim = this.dim;
		return dim;
	}

    public int val(final int x, final int y) {
		Vettore v = this.mat[y];
		return v.val(v);
	}

    public Matrice piu(final Matrice m) {
		if(this.dim != m.dim()) throw new IllegalArgumentException("le matrici devono avere la stessa dimensione per poter eseguire la somma");
		Matrice ris; // variabile di tipo Matrice che conterra' il risultato
		List<Vettore> list = new ArrayList<Vettore>(); // lista di vettori che servira' per creare la matrice 
		for(int y = 0; y < m.dim(); y++) {
			int[] arr = new int[m.dim()]; // array che servira' per creare il vettore
			for(int x = 0; x < m.dim(); x++) {
				arr[x] = this.val(x, y) + m.val(x, y); // somma degli elementi nella stessa posizione delle due matrici operatori. 
			}
			Vettore v = new VettoreDenso(arr); // vettore che conterra' la prima riga della matrice
			list.add(v); // inserimento del vettore nella lista
		}
		ris = new MatriceDensa(list); //  creazione della matrice risultato
		return ris;
	}

    public Matrice per_scalare(final int i) {
		Matrice ris; // variabile di tipo Matrice che conterra' il risultato
		List<Vettore> list = new ArrayList<Vettore>(); // lista di vettori che servira' per creare la matrice 
		for(int y = 0; y < m.dim(); y++) {
			int[] arr = new int[m.dim()]; // array che servira' per creare il vettore
			for(int x = 0; x < m.dim(); x++) {
				arr[x] = this.val(x, y) * i; // moltiplicazione tra elementi della matrice e lo scalare. 
			}
			Vettore v = new VettoreDenso(arr); // vettore che conterra' la prima riga della matrice
			list.add(v); // inserimento del vettore nella lista
		}
		ris = new MatriceDensa(list); //  creazione della matrice risultato
		return ris;
	}

	public Matrice per_matrice(final Matrice m) {	
		if(this.dim != m.dim()) throw new IllegalArgumentException("le matrici devono avere la stessa dimensione per poter eseguire la somma");
		Matrice ris; // matrice risultato
		int val; // variabile che contiene il valore di ris che si sta generando al momento
		List<Vettore> list = new ArrayList(); // lista vettori per creare ris
		for(int yThis = 0; yThis < m.dim(); yThis++) { // spostamento sull'asse y della matrice this.mat
			for(int xM = 0; xM < m.dim(); xM++) { // spostamento sull'asse x della matrice val, in piu indica lo spostamento sull'asse x della matrice m
				int[] arr = new int[m.dim()]; // array per creare il vettore
				for(int yM = 0; yM < m.dim(); yM++) { // spostamento sull'asse y della matrice m
					val += this.val(yM, yThis) * m.val(xM, yM); // moltiplicazione tra gli elementi della matrice per generare val
				}
				arr[xM] = val; // inserimento val nell'array
				val = 0; // reset val
			}
			Vettore v = new VettoreDenso(arr); // creazione vettore
			list.add(v); // inserimento del vettore nella lista
		}
		ris = new MatriceDensa(list); // creazione matrice risultato
		return ris;
	}
	
	public Vettore per_vettore(final Vettore v) {
		return null;
	}
}
