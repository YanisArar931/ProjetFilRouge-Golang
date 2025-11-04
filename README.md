# Mini CRM en Go

Une application en ligne de commande (CLI) pour gérer une liste de contacts :  
ajout, affichage, mise à jour et suppression.

---

## Lancer le projet

### Cloner le dépôt
```bash
git clone https://github.com/YanisArar931/ProjetFilRouge-Golang.git
cd ProjetFilRouge-Golang
```

### Exécuter le programme
```bash
go run .
```

## Fonctionnalités

- Ajouter un contact (nom + email)
- Lister tous les contacts
- Mettre à jour un contact par son ID
- Supprimer un contact par son ID

### Utilisation de Flags
Possibilités d'ajouter un contact directement en ligne de commande dans la console : 
```bash
go run main.go --name="Golang" --email="golang@example.com"
```

### Interface du projet en CLI

-----------------------------
Bienvenue dans notre Mini-CRM !
-----------------------------

1. Ajouter un contact
2. Lister tous les contacts
3. Supprimer un contact par son ID
4. Mettre à jour un contact
5. Quitter l'application

