############################################
###  _____                _ __  __       ###
### |  __ \              | |  \/  |      ###
### | |__) |___  __ _  __| | \  / | ___  ###
### |  _  // _ \/ _` |/ _` | |\/| |/ _ \ ###
### | | \ \  __/ (_| | (_| | |  | |  __/ ###
### |_|  \_\___|\__,_|\__,_|_|  |_|\___| ###
############################################

TODO  (USER)
    #NONEEDED- RUN /reset/ to clear the cache
    - IF YOU MISS SOMETHING, JUST RESTART FROM SCRATCH(please don't try to move to previous/next page with left corner button...)
    - DON'T SPAM ANY SUBMIT BUTTON
    #NOW YOU CAN- DON'T RUN MULTIPLE SCRIPT AT THE SAME TIME(cause it wll cause some data-race issue)
    - DON'T RELOAD ANY PAGE OF THE SCRIPT(cause if you reload page, you'll overwrite data)
    - SCRIPT CAN BE PRETTY SLOW DUE TO SFCC LATENCY(we are looking for more than 100k promotions, campaigns and coupons)

TODO (DEV)
    #NONEEDED - CREATE A /reset/ ENDPOINT TO RESET ALL GLOBAL VARIABLE    ==  safest for zombie proc/variable
	#DONE - ONE ENV(No more front+back only one instance)		  	    ==  faster and more concurrence routine will be usable
																    (anti-data race issue)
	- ??? ADD CONCURRENCE ROUTINE ON DELETE					    ==  faster deleting(not sure - due to thread issue)
	- ??? SOME JS/CSS 										    ==  more user-friendly(but slower)
	- ADD DELETE METHOD ON SFCC ADMIN PANEL                     ==  reach http.delete method api
