public class IntQueue {
  // Campi
  /*- La struttura dati contenente gli elementi dell'IntQueue this. */
  private final int[] elements;

  /*- Gli indici della testa e della coda della IntQueue.
   *  Nello specifico, head indica l'indice del primo elemento di this (-1 se la coda è vuota),
   *  mentre tail indica l'indice della prima posizione disponibile di this (head = tail se la coda è piena).
   */
  private int head, tail;

  /*-
   * ABS FUN:  AF(elements, head, tail)
   *           = [ elements[i] | head <= i < tail ]
   *           = [elements[head], elements[head+1], ..., elements[tail-1]] se -1 < head <= tail
   *           o [elements[head], elements[head+1], ..., elements[elements.size-1], elements[0], ..., elements[tail-1]] se head > tail
   *
   * REP INV:  la coda non contiene più elementi della sua capienza massima,
   *          -1 <= head < size
   *           0 <= tail < size
   *           head == -1 ⇒ tail = 0
   */

  // Costruttori
  /*-
   * Post-condizioni: Inizializza this affinché rappresenti una coda vuota con dimensione massima n.
   *                  Solleva un'eccezione di tipo NegativeArraySizeException se n è negativo.
   */
  public IntQueue(int n) {
    elements = new int[n];
    head = -1;
    tail = 0;

    assert repOK();
  }

  // Metodi
  /*-
   * Effetti collaterali: this è modificato se la coda non è piena
   * Post-condizioni: Aggiunge l'elemento x alla coda this
   *				          e solleva un'eccezione di tipo FullException se la coda è piena
   *                  this_post = this + [x]
   */
  public void enqueue(int x) {
    if (isFull()) throw new FullQueueException("Impossibile aggiungere elemento. Coda piena.");
    if (isEmpty()) head = 0;
    elements[tail] = x;
    tail = (tail + 1) % elements.length;

    assert repOK();
  }

  /*-
   * Effetti collaterali: this è modificato se la coda non è vuota
   * Post-condizioni: Rimuove e restituisce l'elemento in testa alla coda this, se presente.
   *                  e solleva un'eccezione di tipo FullException se la coda è piena
   *                  this = [x1, x2, ..., x_k], k < n, this_post = [x2, ..., x_k]
   */
  public int dequeue() {
    if (isEmpty()) throw new EmptyQueueException("Impossibile estrarre elemento. Coda vuota.");
    int r = elements[head];
    head = (head + 1) % elements.length;
    if (head == tail) {
      head = -1;
      tail = 0;
    }

    assert repOK();

    return r;
  }

  /*-
   * Post-condizioni: restituisce true se la coda this è piena.
   */
  public boolean isFull() {
    return head == tail;
  }

  /*-
   * Post-condizioni: restituisce true se la coda this è vuota.
   */
  public boolean isEmpty() {
    return head == -1;
  }

  /*-
   * Post-condizioni: restituisce true se l'invariante di rappresentazione è valida
   */
  private boolean repOK() {
    return size() <= elements.length
        && elements != null
        && head >= -1
        && head < elements.length
        && tail >= 0
        && tail < elements.length
        && (head != -1 || tail == 0);
  }

  /*-
   * Post-condizioni: restituisce il numero di elementi contenuti in this
   */
  public int size() {
    if (isEmpty()) return 0;
    if (isFull()) return elements.length;
    return (tail - head + elements.length) % elements.length;
  }

  @Override
  public String toString() {
    assert repOK();

    StringBuilder sb = new StringBuilder("IntQueue : [");
    if (!isEmpty()) {
      int i;
      for (i = 0; i < size() - 1; i++)
        sb.append(elements[(head + i) % elements.length]).append(", ");
      sb.append(elements[(head + i) % elements.length]);
    }
    sb.append(']');
    return sb.toString();
  }

  @Override
  public boolean equals(Object obj) {
    if (!(obj instanceof IntQueue)) return false;
    IntQueue other = (IntQueue) obj;

    if (size() != other.size() || elements.length != other.elements.length) return false;
    for (int i = 0; i < size(); i++)
      if (elements[(head + i) % elements.length]
          != other.elements[(other.head + i) % elements.length]) return false;

    return true;
  }
}
