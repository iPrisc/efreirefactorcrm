# Mini-CRM CLI 🚀# Mini-CRM CLI 



## DescriptionUn gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go. Ce projet a été conçu comme un cas pratique pour illustrer les bonnes pratiques de développement Go, incluant :

* Une architecture en packages découplés.

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go. Ce projet illustre les bonnes pratiques de développement Go incluant une architecture découplée, l'injection de dépendances, une CLI professionnelle avec Cobra, et une gestion de configuration externe avec Viper.* L'injection de dépendances via les interfaces.

* La création d'une CLI professionnelle avec Cobra.

## Fonctionnalités ✨* La gestion de configuration externe avec Viper.

* Plusieurs couches de persistance, notamment avec GORM et SQLite.

- **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer

- **Interface en ligne de commande professionnelle** : Commandes et sous-commandes standardisées  ## Contexte du devoir 

- **Configuration externe** : Comportement modifiable sans recompilation

- **Persistance multi-backend** :Nous avons fait évoluer notre application Mini-CRM depuis un simple script jusqu'à un programme modulaire utilisant des packages et la persistance de données via un fichier JSON. L'objectif de ce devoir est de finaliser cette transformation pour en faire un véritable outil en ligne de commande, robuste et configurable.

  - 🧠 **En mémoire** : Stockage éphémère pour les tests

  - 📄 **Fichier JSON** : Sauvegarde simple et lisibleCe devoir est divisé en deux grandes étapes indépendantes mais complémentaires.

  - 🗄️ **Base de données SQLite** : Stockage SQL robuste via GORM

## Partie 1 : Intégration d'une Base de Données avec GORM/SQLite (45%)

## Installation

**Objectif** : Remplacer notre système de stockage JSON par une base de données SQL via l'ORM GORM. Grâce à notre architecture basée sur les interfaces, cette modification devrait se faire sans impacter la logique métier.

### Prérequis

- Go 1.19 ou supérieur### Étapes 



### Compilation1. **Ajouter les dépendances nécessaires**, à la racine du projet ajoutez GORM et son driver SQLite : 

```bash```bash

git clone <votre-repo>    go get gorm.io/gorm

cd mini-crm    go get gorm.io/driver/sqlite

go build -o crm.exe main.go```

```

2. Mettre à jour la struct `Contact`

## Configuration3. Créer le `GORMStore`

4. Implémenter l'interface `Storer`

Le fichier `config.yaml` permet de configurer l'application :5. Intégrer dans `cmd/root.go`



```yaml

# Configuration du Mini-CRM## Partie 2 : Création d'une CLI Professionnelle avec Cobra & Viper (55%)

storage:

  type: "gorm"  # Types: "memory", "json", "gorm"### Étapes

  

  gorm:1. Ajouter les dépendances nécessaire :

    database: "contacts.db"```bash

        go get github.com/spf13/cobra

  json:    go get github.com/spf13/viper

    file: "contacts.json"```

2. Réorganiser les projets : Adoptez une structure de projet orientée Cobra

app:3. Créer le fichier de configuration (`.yaml`) qui permettra de choisir le stockage

  name: "Mini-CRM CLI"4. Implémenter la commande Racine (`cmd/root.go`)

  version: "1.0.0"5. Implémenter les sous-commandes 

```   * Pour chaque fonctionnalité (ajouter, lister, mettre à jour, supprimer), créez un fichier .go dédié dans le package cmd.



## Utilisation## Critères de Réussite



### Aide générale* Le programme **compile et s'exécute sans erreur**.

```bash* Toutes les **sous-commandes** (add, list, update, delete) sont fonctionnelles.

./crm.exe --help* L'application utilise bien la base de données **SQLite** (un fichier .db est créé et mis à jour) lorsque type: "gorm" est configuré dans config.yaml.

```* Il est possible de **basculer sur le stockage json ou memory en modifiant simplement le fichier config.yaml**, sans recompiler.

* Le code est propre, formaté avec gofmt, et raisonnablement commenté.

### Ajouter un contact* Un documentation (readme) claire et complète

```bash

./crm.exe add --name "John Doe" --email "john@example.com"## Fonctionnalités finales attendues

./crm.exe add -n "Jane Smith" -e "jane@example.com"

```* **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contacts.

* **Interface en ligne de commande** : Commandes et sous-commandes claires et standardisées.

### Lister les contacts* **Configuration externe** : Le comportement de l'application (notamment le type de stockage) peut être modifié sans recompiler.

```bash* **Persistance des données** : Support de multiples backends de stockage :

./crm.exe list  * GORM/SQLite : Une base de données SQL robuste contenue dans un simple fichier.

./crm.exe list --format table    # Format tableau (défaut)  * Fichier JSON : Une sauvegarde simple et lisible.

./crm.exe list --format simple   # Format simple  * En mémoire : Un stockage éphémère pour les tests.

```



### Mettre à jour un contact

```bash
./crm.exe update 1 --name "John Smith"
./crm.exe update 1 --email "johnsmith@example.com"
./crm.exe update 1 --name "John Smith" --email "johnsmith@example.com"
```

### Supprimer un contact
```bash
./crm.exe delete 1                # Avec confirmation
./crm.exe delete 1 --force        # Sans confirmation
```

## Changement de backend de stockage

**Sans recompilation !** Modifiez simplement le fichier `config.yaml` :

```yaml
# Pour SQLite (recommandé)
storage:
  type: "gorm"

# Pour JSON
storage:
  type: "json"

# Pour mémoire (tests)
storage:
  type: "memory"
```

## Architecture

### Structure du projet
```
├── cmd/                      # Commandes CLI (Cobra)
│   ├── root.go              # Configuration et commande racine
│   ├── add.go               # Commande d'ajout
│   ├── list.go              # Commande de listage
│   ├── update.go            # Commande de mise à jour
│   └── delete.go            # Commande de suppression
├── internal/storage/         # Couche de stockage
│   ├── storage.go           # Interface commune
│   ├── memory.go            # Backend mémoire
│   ├── json.go              # Backend JSON
│   └── gorm.go              # Backend SQLite/GORM
├── config.yaml              # Configuration
└── main.go                  # Point d'entrée
```

### Interface de stockage
```go
type Storer interface {
    Add(contact *Contact) error
    GetAll() ([]*Contact, error)
    GetByID(id int) (*Contact, error)
    Update(id int, newName, newEmail string) error
    Delete(id int) error
}
```

## Exemples d'usage

```bash
# Ajouter des contacts
./crm.exe add -n "Alice Dupont" -e "alice@example.com"
./crm.exe add -n "Bob Martin" -e "bob@example.com"

# Lister
./crm.exe list

# Mettre à jour
./crm.exe update 1 --name "Alice Bernard"

# Supprimer
./crm.exe delete 2 --force

# Changer de backend vers JSON
# Modifiez config.yaml: type: "json"
./crm.exe list  # Utilise maintenant JSON
```

## Dépendances

- **Cobra** : Framework CLI
- **Viper** : Configuration
- **GORM** : ORM Go
- **SQLite** : Base de données (driver pur Go)

## Tests

```bash
go run test_app.go  # Test des 3 backends
```

---

✅ **Partie 1** : GORM/SQLite intégré  
✅ **Partie 2** : CLI Cobra + Configuration Viper  
✅ **Critères** : Toutes les commandes fonctionnelles, changement de backend sans recompilation