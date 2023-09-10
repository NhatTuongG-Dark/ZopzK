package License

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/Cryptolens/cryptolens-golang/cryptolens"
)


func LicenseGet() bool {
	Key, err := ioutil.ReadFile("build/license.key"); if err != nil {
		log.Println("Failed to Find your License Key File!")
		return false
	}

	token := "WyIzNTExODcxIiwiT2tvYzh1QVEwQkM1SEZNVVdGYkUySWdvMExQWURGRjN0by9DV3o3QiJd"
	publicKey := "<RSAKeyValue><Modulus>3Z53tbLRpe0v7277kEvHP1h1LOVlx1z02p0ApURP/t30+fp1nNgGtwnHWpu8TJFAfyPMCUPXPLo+mXLIlZnha9N/DEf6PdSwj1nGA1IlUTARXh77ndw2Mp4KvEaUBKw4NAbSTX14QIyHca/TED3slKi0grv3OnCqUvvXDtFj56xfHcvnTLhTYHOphob3cRhOICcnouYS02OneKLg42zzYS8+Yuz8F/fUzZA5CyVR4iPLq10tRunBpIw3E/j96CjrGsHLdKPA7rW4X7wWszqvZ5kGYAqWC1ClcpATVQSW4e/wL3RVZrSyx61RWG1o5rm6R4jgwpkfhNVUaw5imAeL1Q==</Modulus><Exponent>AQAB</Exponent></RSAKeyValue>"

	licenseKey, err := cryptolens.KeyActivate(token, cryptolens.KeyActivateArguments{
		ProductId:   11343,
		Key:         string(Key),
	})

	if err != nil || !licenseKey.HasValidSignature(publicKey) {
		return false
	}

	if time.Now().After(licenseKey.Expires) {
		return false
	}

	if licenseKey.F3 {
		log.Println(" [Live Wire DLC is Active on this license]")
		LiveWire = true
	} else {
		LiveWire = false
	}



	return true
}

/*
	for {
		licenseKey, err := cryptolens.KeyActivate(token, cryptolens.KeyActivateArguments{
			ProductId:   11343,
			Key:         string(Key),
		})
	
		if err != nil || !licenseKey.HasValidSignature(publicKey) {
			fmt.Println("License key activation failed!")
			os.Exit(1337)
		}
	
		if time.Now().After(licenseKey.Expires) {
			fmt.Println("License key has expired")
			os.Exit(1337)
		}

		time.Sleep(12 * time.Hour)
	}
	*/