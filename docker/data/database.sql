/*
    This file is used by the docker-compose build command to build the mysql db
    
    Tables:
    * article : stores articles (id, title, header, authors, created_on, updated_on)
    * section : stores sections (id, title, paragraph, position, media, created_on, updated_on,
    parent_id, article_id)
    * comment : stores comments (id, title, content, author, created_on, updated_on)
*/

CREATE TABLE article (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT, 
    title VARCHAR(255),
    header TEXT NOT NULL,
    authors TEXT NOT NULL,
    created_on DATETIME,
    updated_on DATETIME,
    PRIMARY KEY (id)
);

CREATE TABLE section (
     id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
     title VARCHAR(255),
     paragraph TEXT,
     position INT UNSIGNED NOT NULL,
     media TEXT,
     created_on DATETIME,
     updated_on DATETIME,
     parent_id BIGINT UNSIGNED DEFAULT NULL, 
     article_id BIGINT UNSIGNED NOT NULL,
     PRIMARY KEY (id),
     UNIQUE KEY position (position, article_id),
     FOREIGN KEY (article_id) REFERENCES article(id) ON DELETE CASCADE, 
     FOREIGN KEY (parent_id) REFERENCES section(id) ON DELETE CASCADE
);

CREATE TABLE comment (
     id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
     title VARCHAR(255),
     content TEXT,
     author INT UNSIGNED NOT NULL,
     created_on DATETIME,
     updated_on DATETIME,
     PRIMARY KEY (id)
);

/* 
    Default values
*/
INSERT INTO `article` (`id`, `title`, `header`, `authors`, `created_on`, `updated_on`) VALUES
('1', 'Elasmosaurus', 'Elasmosaurus is a genus of plesiosaur that lived in North America during the Campanian stage of the Late Cretaceous period, about 80.5 million years ago. The first specimen was discovered in 1867 near Fort Wallace, Kansas, US, and was sent to the American paleontologist Edward Drinker Cope, who named it E. platyurus in 1868. The generic name means "thin-plate reptile", and the specific name means "flat-tailed". Cope originally reconstructed the skeleton of Elasmosaurus with the skull at the end of the tail, an error which was made light of by the paleontologist Othniel Charles Marsh, and became part of their "Bone Wars" rivalry. Only one incomplete Elasmosaurus skeleton is definitely known, consisting of a fragmentary skull, the spine, and the pectoral and pelvic girdles, and a single species is recognized today; other species are now considered invalid or have been moved to other genera.

Measuring 10.3 meters (34 ft) long, Elasmosaurus would have had a streamlined body with paddle-like limbs, a short tail, a small head, and an extremely long neck. The neck alone was around 7.1 meters (23 ft) long. Along with its relative Albertonectes, it was one of the longest-necked animals to have lived, with the largest number of neck vertebrae known, 72. The skull would have been slender and triangular, with large, fang-like teeth at the front, and smaller teeth towards the back. It had six teeth in each premaxilla of the upper jaw, and may have had 14 teeth in the maxilla and 19 in the dentary of the lower jaw. Most of the neck vertebrae were compressed sideways, and bore a longitudinal crest or keel along the sides.

The family Elasmosauridae was based on the genus Elasmosaurus, the first recognized member of this group of long-necked plesiosaurs. Elasmosaurids were well adapted for aquatic life, and used their flippers for swimming. Contrary to earlier depictions, their necks were not very flexible, and could not be held high above the water surface. It is unknown what their long necks were used for, but they may have had a function in feeding. Elasmosaurids probably ate small fish and marine invertebrates, seizing them with their long teeth, and may have used gastroliths (stomach stones) to help digest their food. Elasmosaurus is known from the Pierre Shale formation, which represents marine deposits from the Western Interior Seaway.', 'AT', '2020-07-22 15:40:14', '2020-07-22 15:40:14'),
('2', 'Tyrannosaurus', 'Tyrannosaurus is a genus of coelurosaurian theropod dinosaurs. The species Tyrannosaurus rex (rex meaning "king" in Latin), often called T. rex or colloquially T-Rex, is one of the most well-represented of the large theropods. Tyrannosaurus lived throughout what is now western North America, on what was then an island continent known as Laramidia. Tyrannosaurus had a much wider range than other tyrannosaurids. Fossils are found in a variety of rock formations dating to the Maastrichtian age of the upper Cretaceous period, 68 to 66 million years ago. It was the last known member of the tyrannosaurids, and among the last non-avian dinosaurs to exist before the Cretaceous–Paleogene extinction event.

Like other tyrannosaurids, Tyrannosaurus was a bipedal carnivore with a massive skull balanced by a long, heavy tail. Relative to its large and powerful hind limbs, Tyrannosaurus forelimbs were short but unusually powerful for their size and had two clawed digits. The most complete specimen measures up to 12.3 meters (40 feet) in length though T. rex could grow to lengths of over 12.3 m (40 ft), up to 3.66 m (12 ft) tall at the hips, and according to most modern estimates 8.4 metric tons (9.3 short tons) to 14 metric tons (15.4 short tons) in weight. Although other theropods rivaled or exceeded Tyrannosaurus rex in size, it is still among the largest known land predators and is estimated to have exerted the strongest bite force among all terrestrial animals. By far the largest carnivore in its environment, Tyrannosaurus rex was most likely an apex predator, preying upon hadrosaurs, armored herbivores like ceratopsians and ankylosaurs, and possibly sauropods. Some experts have suggested the dinosaur was primarily a scavenger. The question of whether Tyrannosaurus was an apex predator or a pure scavenger was among the longest debates in paleontology. Most paleontologists today accept that Tyrannosaurus was both an active predator and a scavenger.

Specimens of Tyrannosaurus rex include some that are nearly complete skeletons. Soft tissue and proteins have been reported in at least one of these specimens. The abundance of fossil material has allowed significant research into many aspects of its biology, including its life history and biomechanics. The feeding habits, physiology and potential speed of Tyrannosaurus rex are a few subjects of debate. Its taxonomy is also controversial, as some scientists consider Tarbosaurus bataar from Asia to be a second Tyrannosaurus species while others maintain Tarbosaurus is a separate genus. Several other genera of North American tyrannosaurids have also been synonymized with Tyrannosaurus.

As the archetypal theropod, Tyrannosaurus has been one of the best-known dinosaurs since the early 20th century, and has been featured in film, advertising, postal stamps, and many other media.', 'TA', '2020-07-22 15:40:37', '2020-07-22 15:40:37'),
('3', 'Branchiosaurus', 'Branchiosaurus (from Greek ''Branchios'', meaning gills and ''Sauros'', meaning lizard) is a genus of small, lightly built early prehistoric amphibians. Fossils have been discovered in strata dating from the late Pennsylvanian Epoch to the Permian Period. The taxa may be invalid; the material referred to the genus may be juvenile specimens of larger amphibians This tiny amphibian was very similar to the Rachitomi, differing primarily in size.[clarification needed] Other distinguishing characteristics include a cartilaginous, less ossified skeleton and a shorter skull. Clear traces of gills are present in many fossilized samples, hence the name.

Originally thought to have vertebrae distinct from rachitomous vertebrae, it was placed in a separate order named Phyllospondyli ("leaf vertebrae"). Later analysis of growth stages showed increasing ossification in larger specimens, which showed that at least some of the species was the larval stage of much larger rachitomes like Eryops, while others represent paedomorphic species which retained the larval gills in adulthood.

Distribution is uncertain, though available fossils come from central Europe, most famous of which are the Permian Niederkirchen Beds around Pfalz, Germany.', 'TAT', '2020-07-22 15:41:11', 
 '2020-07-22 15:41:11');
