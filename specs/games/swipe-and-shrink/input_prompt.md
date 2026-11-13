# Input Prompt & Inspiration

**Date:** 2026-11-13  
**Source:** Riccardo Carlesso (Voice message on Telegram)  
**Game Concept:** Swipe & Shrink (Word Anagram & Depletion Grid)

## Audio Transcription (Verbatim)

> *"Ermete ho trovato questo nuovo gioco in cui tu puoi ruotare le dita e man mano che le lettere si esauriscono poi le parole si poi dopo si cancellano quelle lettere quindi rimani magari con una parola da tre e solo con tre lettere non so se capisci finché quella lettera contribuisce più di una parola c'è se no no crea in respec uno speck.md di questo gioco che domani provo a implementarlo per i bambini in italiano e in inglese e in tedesco e in francese amico mio ok grazie"*

## Key Gameplay Insights from Screenshots

1. **Grid & Swipe**: Bottom tile grid (e.g. 5x3 or 4x4) where players drag/swipe continuously through connected adjacent letters to spell words (like *Wordscapes* / *Scrabble* / *Boggle*).
2. **Target Crossword / Slots**: Top slots display discovered and undiscovered words organized by length (3-letter, 4-letter, 5-letter, 6-letter, 7-letter words).
3. **The "Shrink" / Letter Depletion Mechanic**: 
   - Each tile tracks a reference count of how many remaining target words still require it.
   - When all target words containing a specific letter tile have been discovered, that tile dissolves/pops and disappears from the grid!
   - As the puzzle progresses, the grid empties and simplifies.
   - In the endgame, for the last remaining 3-letter word, only those 3 letters remain on screen, creating an intuitive, satisfying finish.
4. **Target Audience**: Multilingual kids app for Alessandro (8) and Sebastian (6) attending Tandem IMS (supporting **IT**, **EN**, **DE**, **FR**).
