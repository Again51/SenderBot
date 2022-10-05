# SenderBot
## Installation 
### 1ere Étape :
Téléchargez à l'endroit souhaité le dossier.

### 2eme Étape :
- Rendez-vous dans le dossier __cmd__, et ouvrez avec un éditeur de texte le fichier __sender.go__.
- Une fois sender.go ouvert, modifier le chemin du fichier __config.json__ (ligne 28)
![SmugPug](https://i.imgur.com/CMdPLWJ.png)

### 3eme Étape :
- Ouvrez le fichier __config.json__, et ajouter à l'intérieur le Token de votre Bot ainsi que le l'identifiant du canal discord qui sera utilisé.
![SmugPug](https://i.imgur.com/cLoRG9P.png)

### 4eme Étape :
- Depuis votre terminal, rendez-vous dans le dossier sender :
(exemple)
```
cd /home/user/Bureau/sender 
```
- Une fois dans le dossier sender, utiliser la commende suivante :
```
go build -o send
```
- Déplacer send vers usr/bin
```
sudo mv send /usr/bin
```

Et voilà le Bot est enfin prêt à être utilisé !


## Commandes
Envoyer un fichier :
```
send file -n "nom du fichier"
```
```
send file -n "chemin / nom du fichier"
```
Envoyer tous les fichiers contenus dans un dossier :
```
send file all -p "chemin du dossier"
```
Envoyer un fichier sur un autre canal discord :
```
send file -n "nom du fichier" -c "ID du channel"
```

