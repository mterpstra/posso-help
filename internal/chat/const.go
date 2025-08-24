package chat

// Breed of Animal
const ANGUS = "angus"
const NELORE = "nelore"
const BRANGUS = "brangus"
const STA_ZELIA = "sta.zelia"
const CRUZADA = "cruzada"
const CRUZADO = "cruzado"
var BREEDS = []string{ANGUS, NELORE, BRANGUS, STA_ZELIA, CRUZADA, CRUZADO}

// Sex of Animal
const MALE = "m"
const FEMALE = "f"
var SEXES = []string{MALE, FEMALE}

// Dead Types
const MORREU     = "morreu"       // he died
const MORTO      = "morto"        // dead
const NASCEU     = "nasceu morto" // stillborn (born dead)
const ABORTO     = "aborto"       // aborted
const NATIMORTO  = "natimorto"    // stillborn
const NATIMORTOS = "natimortos"   // stillbirths
var DEATHS = []string{MORREU, MORTO, NASCEU, ABORTO, NATIMORTO, NATIMORTOS}

// Pure Breed Specifier
const PURE_BREED = "fft" 
