package tema_2_algebretta

import java.util.Arrays;

/*	Questa classe implemena l'interfaccia Vettore, VettoreDenso rappresenta un vettore che memorizza gli
 *	in un vettore di interi (int).
 *
 *	la classe e' immutabile in quanto non e' possibile eliminare elementi dal vettore, e non si puo'
 *	neanche inserirne di nuovi se non al momento della creazione dell'oggetto (tramite il costruttore).
 *
 *	ABS FUN = ABS(vet, dim) = vet = [el1, el2, ..., eln]
 *							  vet e' un array di interi rappresenta il vettore denso, e dim rappresenta
 *							  la dimensione del vettore.
 *
 *	REP INV = vet != null
 *			  dim >= 0 
 */

public class VettoreDenso implements Vettore {	
	// Attributi
	protected final int[] vet;
	private final int dim;
	
	// Costruttore
	
	/*	Costruttore di VettoreDenso.
	 *
	 *	pre-condizione: arr != null.
	 *	effetti collaterali: potrebbero venir assegnati dei valori a this.vet e this.dim.
	 *	post-condizioni: viene creato l'oggetto di tipo VettoreDenso.
	 *
	 */
	public VettoreDenso(int[] arr) {
		if(arr == null) throw new IllegalArgumentException("L'array e' nullo, quindi non e' possibile creare un vettore partendo da arr.");
		this.vet = arr.clone();
		this.dim = arr.length;
	}

	/*	Costruttore copia di VettoreDenso
	 *
	 *	pre-condizioni: v != null
	 *	post-condizioni: crea un oggeto che e' la copia di other
	 */

	public VettoreDenso(VettoreDenso other) {
		if(other == null) throw new IllegalArgumentException("Impossibile creare una copia di un vettore che ha valore null.");
		int[] arr = other.vet.clone();
		this.vet = arr;
		this.dim = other.dim();
	}

	/*	Costruttore di un vettore nullo
	 *
	 *	pre-condizioni: dim >= 0
	 *	post-condizioni: viene creato un vettore pieno di zeri
	 */

	public VettoreDenso(int dim) {
		if(dim < 0) throw new IllegalArgumentException("Impossibile creare un vettore con dimensione negativa");
		int[] arr = new int[dim];
		for(int i = 0; i < dim; i++) {
			arr[i] = 0;
		}
		this.vet = arr;
		this.dim = dim;
	}

	// Metodi
	public int dim() {
		int dim = this.dim;
		return dim;
	}
	
	int val(final int i) {
		return this.vet[i];
	}

	Vettore per(final int alpha) {
		int[] arr = new int[this.dim]; // creazione array per generare il vettore
		for(int i = 0; i < this.dim; i++) { // ciclo usato per riempire l'array e iterare il vettore this 
			arr[i] = this.vet[i] * alpha; // moltiplicazione dell'i-esima posizione del vettore this per alpha e la metto nell'array
		}
		Vettore v = new VettoreDenso(arr); // creazione del vettore
		return v; 
	}
	
	Vettore piu(final Vettore v) {
		int[] arr = new int[this.dim]; // creazione array per generare il vettore
		for(int i = 0; i < this.dim) { // ciclo usato per iterare i due vettori v e this e per riempire l'array
			arr[i] = this.vet[i] + v.val(i); // somma dei valori di this e v dell'i-esima posizione
		}
		Vettore ris = new VettoreDenso(arr); // creazione vettore
		return ris;
	}
}
