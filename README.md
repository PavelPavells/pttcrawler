pttcrawler
======================

A ptt crawler client to browse broad news and download image if any default image in that article. This tool help you to download those photos for your backup, all the photos still own by original creator. 

- It support multiple platform such as "Windows 8.1" and "MacOS X".


Install
--------------

  go get -u -x github.com/PavelPavells/pttcrawler

Usage
---------------------

  pttcrawler 

All the photos will download to `USERS/Pictures/pttcrawler` and it will separate folder by article name.

For Windows user, it will store in your personal pictures folder.



Options
---------------

- `-w` number of workers. (concurrency), default workers is "25"


Interactive Command
---------------

It support command line interactive command as follow:

- `n`: Display next page aticles.
- `p`: Display previous page articles.
- `o`: Open content folder in finder.
- `d` number: Download article image with specific index, currently it support single index.
- `quit`: Exist current application.

Examples
---------------

Download all photos from Scottie Pippen facebook pages with 10 workers.

        //Run app.
        pttcrawler -w=10
        
        0:[1★][正妹] 勇敢的女孩
        1:[0★](本文已被刪除) [ao3sm345]
        2:[0★](本文已被刪除) [chuhengyi820]
        3:[0★](本文已被刪除) [titan3417]
        4:[8★][公告] 不願上表特 ＆ 優文推薦 ＆ 檢舉建議專區
        ......
        >ptt>

        //Download index [5].
        d 5
        
        //quit application
        quit
     


Snapshot
---------------

![image](snapshot/1.png)

TODOs
---------------

Welcome to file your suggestion in issues.

Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

License
---------------

This package is licensed under MIT license. See LICENSE for details.