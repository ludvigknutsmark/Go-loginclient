var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.post('/getusername', function(req,res,next) {
	var existed;
	var pg = req.pg;
	var fs = req.fs;
	var username = req.body.username;
	const connectionstring = 'X'
	pg.connect(connectionstring, function(err,client,done) {
		if(err) {
			done();
			console.log(err)
			return res.json(err);
		}
	const query = client.query('SELECT * FROM userinfo WHERE username = ($1);',[username]);
	var rows = []
	query.on('row', function(row, result) {
		rows.push(row)
	});

	query.on('end', function(result) {
		existed = result.rowCount
		done();
		return res.json(existed)
	});
	});
	//});

});

router.post('/registeruser', function(req,res,next) {
	var scrypt = req.scrypt;
	var pg = req.pg;
	var fs = req.fs;
	const data = {username: req.body.username, password: req.body.password}
	var key = new Buffer(data.password);
	var scryptParameters = scrypt.paramsSync(0.1);
	/* Hashing the password */
	var kdfResult = scrypt.kdfSync(key, scryptParameters).toString('base64')
	
	const connectionstring = 'X'
	pg.connect(connectionstring, function(err,client,done) {
		if(err) {
			done();
			console.log(err);
			return res.json(err);
		}
	const query = client.query('INSERT INTO userinfo(username, password) VALUES($1, $2)', [data.username, kdfResult], function(err, result) {
		done();
		if(err) {
			return next(err)
		}
		else {
			return res.json("SUCCESS")
		}
	});
	});
	//});
});

router.post('/verifyuser', function(req,res,next){
	var compare = 0
	var results = []
	var scrypt = req.scrypt;
	var pg = req.pg;
	var fs = req.fs;
	const data = {username: req.body.username, password: req.body.password}	
	const connectionstring = 'X'
	pg.connect(connectionstring, function(err,client,done) {
		if(err){
			done();
			console.log(err)
			return res.json(err);
		}
	const query = client.query('SELECT password FROM userinfo WHERE username = ($1);', [data.username]);

	query.on('row', (row)=> {
		results.push(row);
	});

	query.on('end', () => {
		try {
			if(!data.password){
				var err = new Error('test');
				throw err;
			}
			else if(scrypt.verifyKdfSync(new Buffer(results[0]['password'], 'base64'), data.password)) {
				compare = 1
			}
			else {
				compare = 0
			}
		}
		
		catch(err){	
			compare = 0
			return res.json(compare);
		}
		done();
		return res.json(compare)
	});
	});
	//});
});

module.exports = router;
