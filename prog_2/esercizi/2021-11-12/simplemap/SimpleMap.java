/*  Overview: Questa classe implementa una mappa, ovvero una struttura dati in cui per ogni chiave e' associato
 *            un elemento, in piu' le chiavi sono uniche, quindi non e' possibile avere chiavi duplicate.
 *            La classe e' mutabile, siccome e' possibile eseguire operazioni di inserimento, rimozione.
 *            esempio mappa:
 *                          [k1 -> v1]
 *                          [k2 -> v2]
 *                          [k4 -> v3]
 *                          [... -> ...]
 *                          [kk -> vk]                         
 */

public class SimpleMap {
    //Attributi
    
    /*
     *  La struttura dati usata per implementare la mappa e' formata da due liste dove per lo stesso indice
     *  abbiamo nella prima lista la chiave e nella seconda lista il valore.
     */

    List<Integer> values;
    List<String> keys; 
    
    /*
     *  ABS FUN: AF(keys, values)
     *           = keys[i] e values[i] | 0 <= i <= (keys.dim() | values.keys())
     *
     *                  
     *
     *  REP INV: I due array devono avere lo stesso numero di elementi, di conseguenza hanno la stessa dimensione
     *           keys[i] e' unico.
     *           keys != NULL
     *           values != NULL
     *           keys[i] != NULL
     */
    
    //Costruttori
     
    /*
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */

    public SimpleMap() {}
    
    //Metodi
    
    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */

    public void put() {}

    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */
    
    public void remove() {}
    
    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */

    public int get(String k) {}


    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */

    public boolean repOk() {}

    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */

    @Override
    public String toString() {}

    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */
    
    @Override
    public boolean equals() {}

    /*
     *  
     *
     *  Pre-condizioni: 
     *  Effetti collaterali: 
     *  Post-Condizioni:
     */
    
    @Override
    public int hashcode() {}

}























