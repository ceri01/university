package tema_2_algebretta;

/*	Questa classe estende la classe MatriceDensa, e rappresenta una matrice diagonale, ovvero una 
 *	matrice formata da 0 tranne che nella diagonale principale, in questa saranno presenti dei valori qualsiasi.
 *
 *	MatriceDensa e' immutabile, di conseguenza anche questa lo e'.
 *
 * 	Il costruttore di questa classe accetta un array contenente gli elementi che
 * 	saranno presenti nella diagonale principale, questo perche' sarebbe inutile far creare all'utente
 * 	n vettori formati da tutti 0 e un valore qualsiasi, per creare una matrice diagonale, 
 * 	cosi basta inserire la dimensione e un array e verra' fatto in automatico.	
 */

public class MatriceDiagonale extends MatriceDensa {
	// Costruttore
	
	/*	Costruttore di MatriceDiagonale
	 *
	 *	pre-condizioni: arr != null && arr non deve essere vuoto; 
	 *	post-condizioni: viene creato un oggetto istanza di MatriceDiagonale
	 *
	 */
	public MatriceIdentita(int[] arr) {
		if(arr != null || arr.length != 0) throw new IllegalArgumentException("Non e' possibile creare una matrice diagonale partendo da un array vuoto o nullo");
		List<Vettore> v = new ArrayList<Vettore>();
		for(int i = 0; i < dim; i++) {
			int[] varr = new int[dim];
			varr[i] = arr[i];
			v.add(arr);
		}
		super(v);
	}

    /*  Costruttore copia di MatriceDiagonale
     *
     *  per-condizioni: other != null
     *  post-condizioni: crea un oggetto che e' la copia di other
     */

	public MatriceNulla(MatriceNulla other) {
		super(other);
	}
}
