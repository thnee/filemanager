default:

build:
	cd web; just build
	cd api; mv web web_bak
	cd api; cp -r ../web/build ./web
	cd api; just build
	cd api; rm -rf web
	cd api; mv web_bak web
	mv api/main main

clean_build:
	rm -f main
	cd web; rm -rf build

clean_dev:
	cd api; rm -rf tmp
	cd web; rm -rf .svelte-kit

clean:
	just clean_build
	just clean_dev
