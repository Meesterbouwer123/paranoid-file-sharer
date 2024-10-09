# paranoid file sharer
I wanted to fuck around with golang for a while, and got the idea to make a file sharing service in it.
This is mostly a ripoff of https://github.com/Northernside/WiredCloud, I got the idea to make this after finding a few bugs in his implementation.
(also most of the security-focused ideas were from the conversation after he shared his project)

## Security features
~~None yet, I think this program is vulnerable to the most basic of path traversal exploits.~~ I added some mitigations, it should now be a tiny bit more secure.

I do plan to add a bunch of these features in the future:
- [X] Server-side encryption to mitigate path traversal attacks (the program tries to decrypt it first, so you can't get the normal file contents)
- [ ] Client-side encryption to mitigate malicious admins
- [ ] Server-side backup for the above in case the client has Javascript enabled + the appropriate opsec warnings

Also please let me know if you have any other features i should add :D
