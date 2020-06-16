
package openimmo

import "encoding/xml"

// Openimmo was generated 2020-06-17 00:05:40
type Openimmo struct {
	XMLName      xml.Name `xml:"openimmo" json:"openimmo,omitempty"`
	Text         string   `xml:",chardata" json:"text,omitempty"`
	Uebertragung struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		Art            string `xml:"art,attr" json:"art,omitempty"`
		Umfang         string `xml:"umfang,attr" json:"umfang,omitempty"`
		Modus          string `xml:"modus,attr" json:"modus,omitempty"`
		Version        string `xml:"version,attr" json:"version,omitempty"`
		Sendersoftware string `xml:"sendersoftware,attr" json:"sendersoftware,omitempty"`
		Senderversion  string `xml:"senderversion,attr" json:"senderversion,omitempty"`
		TechnEmail     string `xml:"techn_email,attr" json:"techn_email,omitempty"`
		Timestamp      string `xml:"timestamp,attr" json:"timestamp,omitempty"`
		RegiID         string `xml:"regi_id,attr" json:"regi_id,omitempty"`
	} `xml:"uebertragung" json:"uebertragung,omitempty"`
	Anbieter struct {
		Text       string `xml:",chardata" json:"text,omitempty"`
		Anbieternr struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"anbieternr" json:"anbieternr,omitempty"`
		Firma struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"firma" json:"firma,omitempty"`
		OpenimmoAnid struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"openimmo_anid" json:"openimmo_anid,omitempty"`
		Lizenzkennung struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"lizenzkennung" json:"lizenzkennung,omitempty"`
		Anhang struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			Location    string `xml:"location,attr" json:"location,omitempty"`
			Gruppe      string `xml:"gruppe,attr" json:"gruppe,omitempty"`
			Anhangtitel struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"anhangtitel" json:"anhangtitel,omitempty"`
			Format struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"format" json:"format,omitempty"`
			Daten struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Pfad struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"pfad" json:"pfad,omitempty"`
			} `xml:"daten" json:"daten,omitempty"`
		} `xml:"anhang" json:"anhang,omitempty"`
		Immobilie struct {
			Text            string `xml:",chardata" json:"text,omitempty"`
			Objektkategorie struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				Nutzungsart struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					WOHNEN  string `xml:"WOHNEN,attr" json:"wohnen,omitempty"`
					GEWERBE string `xml:"GEWERBE,attr" json:"gewerbe,omitempty"`
					ANLAGE  string `xml:"ANLAGE,attr" json:"anlage,omitempty"`
					WAZ     string `xml:"WAZ,attr" json:"waz,omitempty"`
				} `xml:"nutzungsart" json:"nutzungsart,omitempty"`
				Vermarktungsart struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					KAUF       string `xml:"KAUF,attr" json:"kauf,omitempty"`
					MIETEPACHT string `xml:"MIETE_PACHT,attr" json:"miete_pacht,omitempty"`
					ERBPACHT   string `xml:"ERBPACHT,attr" json:"erbpacht,omitempty"`
					LEASING    string `xml:"LEASING,attr" json:"leasing,omitempty"`
				} `xml:"vermarktungsart" json:"vermarktungsart,omitempty"`
				Objektart struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					Zimmer struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						Zimmertyp string `xml:"zimmertyp,attr" json:"zimmertyp,omitempty"`
					} `xml:"zimmer" json:"zimmer,omitempty"`
					ObjektartZusatz struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"objektart_zusatz" json:"objektart_zusatz,omitempty"`
				} `xml:"objektart" json:"objektart,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
			} `xml:"objektkategorie" json:"objektkategorie,omitempty"`
			Geo struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Plz  struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"plz" json:"plz,omitempty"`
				Ort struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ort" json:"ort,omitempty"`
				Geokoordinaten struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					Breitengrad string `xml:"breitengrad,attr" json:"breitengrad,omitempty"`
					Laengengrad string `xml:"laengengrad,attr" json:"laengengrad,omitempty"`
				} `xml:"geokoordinaten" json:"geokoordinaten,omitempty"`
				Strasse struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"strasse" json:"strasse,omitempty"`
				Hausnummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hausnummer" json:"hausnummer,omitempty"`
				Bundesland struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bundesland" json:"bundesland,omitempty"`
				Land struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					IsoLand string `xml:"iso_land,attr" json:"iso_land,omitempty"`
				} `xml:"land" json:"land,omitempty"`
				Gemeindecode struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gemeindecode" json:"gemeindecode,omitempty"`
				Flur struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"flur" json:"flur,omitempty"`
				Flurstueck struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"flurstueck" json:"flurstueck,omitempty"`
				Gemarkung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gemarkung" json:"gemarkung,omitempty"`
				Etage struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"etage" json:"etage,omitempty"`
				AnzahlEtagen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_etagen" json:"anzahl_etagen,omitempty"`
				LageImBau struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					LINKS  string `xml:"LINKS,attr" json:"links,omitempty"`
					RECHTS string `xml:"RECHTS,attr" json:"rechts,omitempty"`
					VORNE  string `xml:"VORNE,attr" json:"vorne,omitempty"`
					HINTEN string `xml:"HINTEN,attr" json:"hinten,omitempty"`
				} `xml:"lage_im_bau" json:"lage_im_bau,omitempty"`
				Wohnungsnr struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wohnungsnr" json:"wohnungsnr,omitempty"`
				LageGebiet struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					Gebiete string `xml:"gebiete,attr" json:"gebiete,omitempty"`
				} `xml:"lage_gebiet" json:"lage_gebiet,omitempty"`
				RegionalerZusatz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"regionaler_zusatz" json:"regionaler_zusatz,omitempty"`
				KartenMakro struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"karten_makro" json:"karten_makro,omitempty"`
				KartenMikro struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"karten_mikro" json:"karten_mikro,omitempty"`
				Virtuelletour struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"virtuelletour" json:"virtuelletour,omitempty"`
				Luftbildern struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"luftbildern" json:"luftbildern,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"geo" json:"geo,omitempty"`
			Kontaktperson struct {
				Text          string `xml:",chardata" json:"text,omitempty"`
				EmailZentrale struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_zentrale" json:"email_zentrale,omitempty"`
				EmailDirekt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_direkt" json:"email_direkt,omitempty"`
				TelZentrale struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_zentrale" json:"tel_zentrale,omitempty"`
				TelDurchw struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_durchw" json:"tel_durchw,omitempty"`
				TelFax struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_fax" json:"tel_fax,omitempty"`
				TelHandy struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_handy" json:"tel_handy,omitempty"`
				Name struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"name" json:"name,omitempty"`
				Vorname struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"vorname" json:"vorname,omitempty"`
				Titel struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"titel" json:"titel,omitempty"`
				Anrede struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anrede" json:"anrede,omitempty"`
				AnredeBrief struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anrede_brief" json:"anrede_brief,omitempty"`
				Firma struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"firma" json:"firma,omitempty"`
				Zusatzfeld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zusatzfeld" json:"zusatzfeld,omitempty"`
				Strasse struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"strasse" json:"strasse,omitempty"`
				Hausnummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hausnummer" json:"hausnummer,omitempty"`
				Plz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"plz" json:"plz,omitempty"`
				Ort struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ort" json:"ort,omitempty"`
				Postfach struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postfach" json:"postfach,omitempty"`
				PostfPlz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postf_plz" json:"postf_plz,omitempty"`
				PostfOrt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postf_ort" json:"postf_ort,omitempty"`
				Land struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					IsoLand string `xml:"iso_land,attr" json:"iso_land,omitempty"`
				} `xml:"land" json:"land,omitempty"`
				EmailPrivat struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_privat" json:"email_privat,omitempty"`
				EmailSonstige struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					Emailart  string `xml:"emailart,attr" json:"emailart,omitempty"`
					Bemerkung string `xml:"bemerkung,attr" json:"bemerkung,omitempty"`
				} `xml:"email_sonstige" json:"email_sonstige,omitempty"`
				EmailFeedback struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_feedback" json:"email_feedback,omitempty"`
				TelPrivat struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_privat" json:"tel_privat,omitempty"`
				TelSonstige struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Telefonart string `xml:"telefonart,attr" json:"telefonart,omitempty"`
					Bemerkung  string `xml:"bemerkung,attr" json:"bemerkung,omitempty"`
				} `xml:"tel_sonstige" json:"tel_sonstige,omitempty"`
				URL struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"url" json:"url,omitempty"`
				Adressfreigabe struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"adressfreigabe" json:"adressfreigabe,omitempty"`
				Personennummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"personennummer" json:"personennummer,omitempty"`
				Immobilientreuhaenderid struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"immobilientreuhaenderid" json:"immobilientreuhaenderid,omitempty"`
				Foto struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Location string `xml:"location,attr" json:"location,omitempty"`
					Format   struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"format" json:"format,omitempty"`
					Daten struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Pfad struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"pfad" json:"pfad,omitempty"`
					} `xml:"daten" json:"daten,omitempty"`
				} `xml:"foto" json:"foto,omitempty"`
				Freitextfeld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"freitextfeld" json:"freitextfeld,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"kontaktperson" json:"kontaktperson,omitempty"`
			WeitereAdresse struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Adressart string `xml:"adressart,attr" json:"adressart,omitempty"`
				Vorname   struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"vorname" json:"vorname,omitempty"`
				Name struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"name" json:"name,omitempty"`
				Titel struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"titel" json:"titel,omitempty"`
				Anrede struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anrede" json:"anrede,omitempty"`
				AnredeBrief struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anrede_brief" json:"anrede_brief,omitempty"`
				Firma struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"firma" json:"firma,omitempty"`
				Zusatzfeld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zusatzfeld" json:"zusatzfeld,omitempty"`
				Strasse struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"strasse" json:"strasse,omitempty"`
				Hausnummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hausnummer" json:"hausnummer,omitempty"`
				Plz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"plz" json:"plz,omitempty"`
				Ort struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ort" json:"ort,omitempty"`
				Postfach struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postfach" json:"postfach,omitempty"`
				PostfPlz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postf_plz" json:"postf_plz,omitempty"`
				PostfOrt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"postf_ort" json:"postf_ort,omitempty"`
				Land struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					IsoLand string `xml:"iso_land,attr" json:"iso_land,omitempty"`
				} `xml:"land" json:"land,omitempty"`
				EmailZentrale struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_zentrale" json:"email_zentrale,omitempty"`
				EmailDirekt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_direkt" json:"email_direkt,omitempty"`
				EmailPrivat struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"email_privat" json:"email_privat,omitempty"`
				EmailSonstige struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					Emailart  string `xml:"emailart,attr" json:"emailart,omitempty"`
					Bemerkung string `xml:"bemerkung,attr" json:"bemerkung,omitempty"`
				} `xml:"email_sonstige" json:"email_sonstige,omitempty"`
				TelDurchw struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_durchw" json:"tel_durchw,omitempty"`
				TelZentrale struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_zentrale" json:"tel_zentrale,omitempty"`
				TelHandy struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_handy" json:"tel_handy,omitempty"`
				TelFax struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_fax" json:"tel_fax,omitempty"`
				TelPrivat struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"tel_privat" json:"tel_privat,omitempty"`
				TelSonstige struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Telefonart string `xml:"telefonart,attr" json:"telefonart,omitempty"`
					Bemerkung  string `xml:"bemerkung,attr" json:"bemerkung,omitempty"`
				} `xml:"tel_sonstige" json:"tel_sonstige,omitempty"`
				URL struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"url" json:"url,omitempty"`
				Adressfreigabe struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"adressfreigabe" json:"adressfreigabe,omitempty"`
				Personennummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"personennummer" json:"personennummer,omitempty"`
				Freitextfeld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"freitextfeld" json:"freitextfeld,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"weitere_adresse" json:"weitere_adresse,omitempty"`
			Preise struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Kaufpreis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaufpreis" json:"kaufpreis,omitempty"`
				Kaufpreisnetto struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaufpreisnetto" json:"kaufpreisnetto,omitempty"`
				Kaufpreisbrutto struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaufpreisbrutto" json:"kaufpreisbrutto,omitempty"`
				Nettokaltmiete struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nettokaltmiete" json:"nettokaltmiete,omitempty"`
				Kaltmiete struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaltmiete" json:"kaltmiete,omitempty"`
				Warmmiete struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"warmmiete" json:"warmmiete,omitempty"`
				Nebenkosten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nebenkosten" json:"nebenkosten,omitempty"`
				HeizkostenEnthalten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"heizkosten_enthalten" json:"heizkosten_enthalten,omitempty"`
				Heizkosten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"heizkosten" json:"heizkosten,omitempty"`
				ZzgMehrwertsteuer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zzg_mehrwertsteuer" json:"zzg_mehrwertsteuer,omitempty"`
				Mietzuschlaege struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"mietzuschlaege" json:"mietzuschlaege,omitempty"`
				Hauptmietzinsnetto struct {
					Text             string `xml:",chardata" json:"text,omitempty"`
					Hauptmietzinsust string `xml:"hauptmietzinsust,attr" json:"hauptmietzinsust,omitempty"`
				} `xml:"hauptmietzinsnetto" json:"hauptmietzinsnetto,omitempty"`
				Pauschalmiete struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"pauschalmiete" json:"pauschalmiete,omitempty"`
				Betriebskostennetto struct {
					Text              string `xml:",chardata" json:"text,omitempty"`
					Betriebskostenust string `xml:"betriebskostenust,attr" json:"betriebskostenust,omitempty"`
				} `xml:"betriebskostennetto" json:"betriebskostennetto,omitempty"`
				Pacht struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"pacht" json:"pacht,omitempty"`
				Erbpacht struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"erbpacht" json:"erbpacht,omitempty"`
				Hausgeld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hausgeld" json:"hausgeld,omitempty"`
				Abstand struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"abstand" json:"abstand,omitempty"`
				PreisZeitraumVon struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"preis_zeitraum_von" json:"preis_zeitraum_von,omitempty"`
				PreisZeitraumBis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"preis_zeitraum_bis" json:"preis_zeitraum_bis,omitempty"`
				PreisZeiteinheit struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					Zeiteinheit string `xml:"zeiteinheit,attr" json:"zeiteinheit,omitempty"`
				} `xml:"preis_zeiteinheit" json:"preis_zeiteinheit,omitempty"`
				MietpreisProQm struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"mietpreis_pro_qm" json:"mietpreis_pro_qm,omitempty"`
				KaufpreisProQm struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaufpreis_pro_qm" json:"kaufpreis_pro_qm,omitempty"`
				Provisionspflichtig struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"provisionspflichtig" json:"provisionspflichtig,omitempty"`
				ProvisionTeilen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"provision_teilen" json:"provision_teilen,omitempty"`
				InnenCourtage struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					MitMwst string `xml:"mit_mwst,attr" json:"mit_mwst,omitempty"`
				} `xml:"innen_courtage" json:"innen_courtage,omitempty"`
				AussenCourtage struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					MitMwst string `xml:"mit_mwst,attr" json:"mit_mwst,omitempty"`
				} `xml:"aussen_courtage" json:"aussen_courtage,omitempty"`
				CourtageHinweis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"courtage_hinweis" json:"courtage_hinweis,omitempty"`
				Provisionnetto struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"provisionnetto" json:"provisionnetto,omitempty"`
				Provisionbrutto struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"provisionbrutto" json:"provisionbrutto,omitempty"`
				Waehrung struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					IsoWaehrung string `xml:"iso_waehrung,attr" json:"iso_waehrung,omitempty"`
				} `xml:"waehrung" json:"waehrung,omitempty"`
				MwstSatz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"mwst_satz" json:"mwst_satz,omitempty"`
				MwstGesamt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"mwst_gesamt" json:"mwst_gesamt,omitempty"`
				FreitextPreis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"freitext_preis" json:"freitext_preis,omitempty"`
				XFache struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"x_fache" json:"x_fache,omitempty"`
				Nettorendite struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nettorendite" json:"nettorendite,omitempty"`
				NettorenditeSoll struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nettorendite_soll" json:"nettorendite_soll,omitempty"`
				NettorenditeIst struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nettorendite_ist" json:"nettorendite_ist,omitempty"`
				MieteinnahmenIst struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					Periode string `xml:"periode,attr" json:"periode,omitempty"`
				} `xml:"mieteinnahmen_ist" json:"mieteinnahmen_ist,omitempty"`
				MieteinnahmenSoll struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					Periode string `xml:"periode,attr" json:"periode,omitempty"`
				} `xml:"mieteinnahmen_soll" json:"mieteinnahmen_soll,omitempty"`
				Erschliessungskosten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"erschliessungskosten" json:"erschliessungskosten,omitempty"`
				Kaution struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaution" json:"kaution,omitempty"`
				KautionText struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kaution_text" json:"kaution_text,omitempty"`
				Geschaeftsguthaben struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"geschaeftsguthaben" json:"geschaeftsguthaben,omitempty"`
				StpCarport struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_carport" json:"stp_carport,omitempty"`
				StpDuplex struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_duplex" json:"stp_duplex,omitempty"`
				StpFreiplatz struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_freiplatz" json:"stp_freiplatz,omitempty"`
				StpGarage struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_garage" json:"stp_garage,omitempty"`
				StpParkhaus struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_parkhaus" json:"stp_parkhaus,omitempty"`
				StpTiefgarage struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
				} `xml:"stp_tiefgarage" json:"stp_tiefgarage,omitempty"`
				StpSonstige struct {
					Text                string `xml:",chardata" json:"text,omitempty"`
					Stellplatzmiete     string `xml:"stellplatzmiete,attr" json:"stellplatzmiete,omitempty"`
					Stellplatzkaufpreis string `xml:"stellplatzkaufpreis,attr" json:"stellplatzkaufpreis,omitempty"`
					Anzahl              string `xml:"anzahl,attr" json:"anzahl,omitempty"`
					Platzart            string `xml:"platzart,attr" json:"platzart,omitempty"`
					Bemerkung           string `xml:"bemerkung,attr" json:"bemerkung,omitempty"`
				} `xml:"stp_sonstige" json:"stp_sonstige,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Feld struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"name" json:"name,omitempty"`
						Wert struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"wert" json:"wert,omitempty"`
						Typ struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"typ" json:"typ,omitempty"`
						Modus struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"modus" json:"modus,omitempty"`
					} `xml:"feld" json:"feld,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"preise" json:"preise,omitempty"`
			Bieterverfahren struct {
				Text                string `xml:",chardata" json:"text,omitempty"`
				BeginnAngebotsphase struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"beginn_angebotsphase" json:"beginn_angebotsphase,omitempty"`
				Besichtigungstermin struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"besichtigungstermin" json:"besichtigungstermin,omitempty"`
				HoechstgebotZeigen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hoechstgebot_zeigen" json:"hoechstgebot_zeigen,omitempty"`
				Mindestpreis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"mindestpreis" json:"mindestpreis,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Feld struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"name" json:"name,omitempty"`
						Wert struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"wert" json:"wert,omitempty"`
						Typ struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"typ" json:"typ,omitempty"`
						Modus struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"modus" json:"modus,omitempty"`
					} `xml:"feld" json:"feld,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"bieterverfahren" json:"bieterverfahren,omitempty"`
			Versteigerung struct {
				Text                string `xml:",chardata" json:"text,omitempty"`
				Zwangsversteigerung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zwangsversteigerung" json:"zwangsversteigerung,omitempty"`
				Aktenzeichen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"aktenzeichen" json:"aktenzeichen,omitempty"`
				Zvtermin struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zvtermin" json:"zvtermin,omitempty"`
				Zusatztermin struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zusatztermin" json:"zusatztermin,omitempty"`
				Amtsgericht struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"amtsgericht" json:"amtsgericht,omitempty"`
				Verkehrswert struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"verkehrswert" json:"verkehrswert,omitempty"`
			} `xml:"versteigerung" json:"versteigerung,omitempty"`
			Flaechen struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				Wohnflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wohnflaeche" json:"wohnflaeche,omitempty"`
				Nutzflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nutzflaeche" json:"nutzflaeche,omitempty"`
				Gesamtflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gesamtflaeche" json:"gesamtflaeche,omitempty"`
				Ladenflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ladenflaeche" json:"ladenflaeche,omitempty"`
				Lagerflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"lagerflaeche" json:"lagerflaeche,omitempty"`
				Verkaufsflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"verkaufsflaeche" json:"verkaufsflaeche,omitempty"`
				Freiflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"freiflaeche" json:"freiflaeche,omitempty"`
				Bueroflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bueroflaeche" json:"bueroflaeche,omitempty"`
				Bueroteilflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bueroteilflaeche" json:"bueroteilflaeche,omitempty"`
				Fensterfront struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"fensterfront" json:"fensterfront,omitempty"`
				Verwaltungsflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"verwaltungsflaeche" json:"verwaltungsflaeche,omitempty"`
				Gastroflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gastroflaeche" json:"gastroflaeche,omitempty"`
				Grz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"grz" json:"grz,omitempty"`
				Gfz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gfz" json:"gfz,omitempty"`
				Bmz struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bmz" json:"bmz,omitempty"`
				Bgf struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bgf" json:"bgf,omitempty"`
				Grundstuecksflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"grundstuecksflaeche" json:"grundstuecksflaeche,omitempty"`
				Sonstflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"sonstflaeche" json:"sonstflaeche,omitempty"`
				AnzahlZimmer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_zimmer" json:"anzahl_zimmer,omitempty"`
				AnzahlSchlafzimmer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_schlafzimmer" json:"anzahl_schlafzimmer,omitempty"`
				AnzahlBadezimmer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_badezimmer" json:"anzahl_badezimmer,omitempty"`
				AnzahlSepWc struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_sep_wc" json:"anzahl_sep_wc,omitempty"`
				AnzahlBalkone struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_balkone" json:"anzahl_balkone,omitempty"`
				AnzahlTerrassen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_terrassen" json:"anzahl_terrassen,omitempty"`
				AnzahlLogia struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_logia" json:"anzahl_logia,omitempty"`
				BalkonTerrasseFlaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"balkon_terrasse_flaeche" json:"balkon_terrasse_flaeche,omitempty"`
				AnzahlWohnSchlafzimmer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_wohn_schlafzimmer" json:"anzahl_wohn_schlafzimmer,omitempty"`
				Gartenflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gartenflaeche" json:"gartenflaeche,omitempty"`
				Kellerflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kellerflaeche" json:"kellerflaeche,omitempty"`
				FensterfrontQm struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"fensterfront_qm" json:"fensterfront_qm,omitempty"`
				Grundstuecksfront struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"grundstuecksfront" json:"grundstuecksfront,omitempty"`
				Dachbodenflaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"dachbodenflaeche" json:"dachbodenflaeche,omitempty"`
				TeilbarAb struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"teilbar_ab" json:"teilbar_ab,omitempty"`
				BeheizbareFlaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"beheizbare_flaeche" json:"beheizbare_flaeche,omitempty"`
				AnzahlStellplaetze struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_stellplaetze" json:"anzahl_stellplaetze,omitempty"`
				PlaetzeGastraum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"plaetze_gastraum" json:"plaetze_gastraum,omitempty"`
				AnzahlBetten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_betten" json:"anzahl_betten,omitempty"`
				AnzahlTagungsraeume struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_tagungsraeume" json:"anzahl_tagungsraeume,omitempty"`
				VermietbareFlaeche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"vermietbare_flaeche" json:"vermietbare_flaeche,omitempty"`
				AnzahlWohneinheiten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_wohneinheiten" json:"anzahl_wohneinheiten,omitempty"`
				AnzahlGewerbeeinheiten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"anzahl_gewerbeeinheiten" json:"anzahl_gewerbeeinheiten,omitempty"`
				Einliegerwohnung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"einliegerwohnung" json:"einliegerwohnung,omitempty"`
				Kubatur struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kubatur" json:"kubatur,omitempty"`
				Ausnuetzungsziffer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ausnuetzungsziffer" json:"ausnuetzungsziffer,omitempty"`
				Flaechevon struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"flaechevon" json:"flaechevon,omitempty"`
				Flaechebis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"flaechebis" json:"flaechebis,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Feld struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"name" json:"name,omitempty"`
						Wert struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"wert" json:"wert,omitempty"`
						Typ struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"typ" json:"typ,omitempty"`
						Modus struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"modus" json:"modus,omitempty"`
					} `xml:"feld" json:"feld,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"flaechen" json:"flaechen,omitempty"`
			Ausstattung struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				AusstattKategorie struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ausstatt_kategorie" json:"ausstatt_kategorie,omitempty"`
				WgGeeignet struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wg_geeignet" json:"wg_geeignet,omitempty"`
				RaeumeVeraenderbar struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"raeume_veraenderbar" json:"raeume_veraenderbar,omitempty"`
				Bad struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					DUSCHE  string `xml:"DUSCHE,attr" json:"dusche,omitempty"`
					WANNE   string `xml:"WANNE,attr" json:"wanne,omitempty"`
					FENSTER string `xml:"FENSTER,attr" json:"fenster,omitempty"`
					BIDET   string `xml:"BIDET,attr" json:"bidet,omitempty"`
					PISSOIR string `xml:"PISSOIR,attr" json:"pissoir,omitempty"`
				} `xml:"bad" json:"bad,omitempty"`
				Kueche struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					EBK    string `xml:"EBK,attr" json:"ebk,omitempty"`
					OFFEN  string `xml:"OFFEN,attr" json:"offen,omitempty"`
					PANTRY string `xml:"PANTRY,attr" json:"pantry,omitempty"`
				} `xml:"kueche" json:"kueche,omitempty"`
				Boden struct {
					Text          string `xml:",chardata" json:"text,omitempty"`
					FLIESEN       string `xml:"FLIESEN,attr" json:"fliesen,omitempty"`
					STEIN         string `xml:"STEIN,attr" json:"stein,omitempty"`
					TEPPICH       string `xml:"TEPPICH,attr" json:"teppich,omitempty"`
					PARKETT       string `xml:"PARKETT,attr" json:"parkett,omitempty"`
					FERTIGPARKETT string `xml:"FERTIGPARKETT,attr" json:"fertigparkett,omitempty"`
					LAMINAT       string `xml:"LAMINAT,attr" json:"laminat,omitempty"`
					DIELEN        string `xml:"DIELEN,attr" json:"dielen,omitempty"`
					KUNSTSTOFF    string `xml:"KUNSTSTOFF,attr" json:"kunststoff,omitempty"`
					ESTRICH       string `xml:"ESTRICH,attr" json:"estrich,omitempty"`
					DOPPELBODEN   string `xml:"DOPPELBODEN,attr" json:"doppelboden,omitempty"`
					LINOLEUM      string `xml:"LINOLEUM,attr" json:"linoleum,omitempty"`
					MARMOR        string `xml:"MARMOR,attr" json:"marmor,omitempty"`
					TERRAKOTTA    string `xml:"TERRAKOTTA,attr" json:"terrakotta,omitempty"`
					GRANIT        string `xml:"GRANIT,attr" json:"granit,omitempty"`
				} `xml:"boden" json:"boden,omitempty"`
				Kamin struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kamin" json:"kamin,omitempty"`
				Heizungsart struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					OFEN      string `xml:"OFEN,attr" json:"ofen,omitempty"`
					ETAGE     string `xml:"ETAGE,attr" json:"etage,omitempty"`
					ZENTRAL   string `xml:"ZENTRAL,attr" json:"zentral,omitempty"`
					FERN      string `xml:"FERN,attr" json:"fern,omitempty"`
					FUSSBODEN string `xml:"FUSSBODEN,attr" json:"fussboden,omitempty"`
				} `xml:"heizungsart" json:"heizungsart,omitempty"`
				Befeuerung struct {
					Text          string `xml:",chardata" json:"text,omitempty"`
					OEL           string `xml:"OEL,attr" json:"oel,omitempty"`
					GAS           string `xml:"GAS,attr" json:"gas,omitempty"`
					ELEKTRO       string `xml:"ELEKTRO,attr" json:"elektro,omitempty"`
					ALTERNATIV    string `xml:"ALTERNATIV,attr" json:"alternativ,omitempty"`
					SOLAR         string `xml:"SOLAR,attr" json:"solar,omitempty"`
					ERDWAERME     string `xml:"ERDWAERME,attr" json:"erdwaerme,omitempty"`
					LUFTWP        string `xml:"LUFTWP,attr" json:"luftwp,omitempty"`
					FERN          string `xml:"FERN,attr" json:"fern,omitempty"`
					BLOCK         string `xml:"BLOCK,attr" json:"block,omitempty"`
					WASSERELEKTRO string `xml:"WASSER-ELEKTRO,attr" json:"wasser-elektro,omitempty"`
					PELLET        string `xml:"PELLET,attr" json:"pellet,omitempty"`
				} `xml:"befeuerung" json:"befeuerung,omitempty"`
				Klimatisiert struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"klimatisiert" json:"klimatisiert,omitempty"`
				Fahrstuhl struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					PERSONEN string `xml:"PERSONEN,attr" json:"personen,omitempty"`
					LASTEN   string `xml:"LASTEN,attr" json:"lasten,omitempty"`
				} `xml:"fahrstuhl" json:"fahrstuhl,omitempty"`
				Stellplatzart struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					GARAGE     string `xml:"GARAGE,attr" json:"garage,omitempty"`
					TIEFGARAGE string `xml:"TIEFGARAGE,attr" json:"tiefgarage,omitempty"`
					CARPORT    string `xml:"CARPORT,attr" json:"carport,omitempty"`
					FREIPLATZ  string `xml:"FREIPLATZ,attr" json:"freiplatz,omitempty"`
					PARKHAUS   string `xml:"PARKHAUS,attr" json:"parkhaus,omitempty"`
					DUPLEX     string `xml:"DUPLEX,attr" json:"duplex,omitempty"`
				} `xml:"stellplatzart" json:"stellplatzart,omitempty"`
				Gartennutzung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gartennutzung" json:"gartennutzung,omitempty"`
				AusrichtBalkonTerrasse struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					NORD     string `xml:"NORD,attr" json:"nord,omitempty"`
					OST      string `xml:"OST,attr" json:"ost,omitempty"`
					SUED     string `xml:"SUED,attr" json:"sued,omitempty"`
					WEST     string `xml:"WEST,attr" json:"west,omitempty"`
					NORDOST  string `xml:"NORDOST,attr" json:"nordost,omitempty"`
					NORDWEST string `xml:"NORDWEST,attr" json:"nordwest,omitempty"`
					SUEDOST  string `xml:"SUEDOST,attr" json:"suedost,omitempty"`
					SUEDWEST string `xml:"SUEDWEST,attr" json:"suedwest,omitempty"`
				} `xml:"ausricht_balkon_terrasse" json:"ausricht_balkon_terrasse,omitempty"`
				Moebliert struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Moeb string `xml:"moeb,attr" json:"moeb,omitempty"`
				} `xml:"moebliert" json:"moebliert,omitempty"`
				Rollstuhlgerecht struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"rollstuhlgerecht" json:"rollstuhlgerecht,omitempty"`
				KabelSatTv struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kabel_sat_tv" json:"kabel_sat_tv,omitempty"`
				Barrierefrei struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"barrierefrei" json:"barrierefrei,omitempty"`
				Sauna struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"sauna" json:"sauna,omitempty"`
				Swimmingpool struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"swimmingpool" json:"swimmingpool,omitempty"`
				WaschTrockenraum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wasch_trockenraum" json:"wasch_trockenraum,omitempty"`
				Wintergarten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wintergarten" json:"wintergarten,omitempty"`
				DvVerkabelung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"dv_verkabelung" json:"dv_verkabelung,omitempty"`
				Rampe struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"rampe" json:"rampe,omitempty"`
				Hebebuehne struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hebebuehne" json:"hebebuehne,omitempty"`
				Kran struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kran" json:"kran,omitempty"`
				Gastterrasse struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gastterrasse" json:"gastterrasse,omitempty"`
				Stromanschlusswert struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"stromanschlusswert" json:"stromanschlusswert,omitempty"`
				KantineCafeteria struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kantine_cafeteria" json:"kantine_cafeteria,omitempty"`
				Teekueche struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"teekueche" json:"teekueche,omitempty"`
				Hallenhoehe struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hallenhoehe" json:"hallenhoehe,omitempty"`
				AngeschlGastronomie struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					HOTELRESTAURANT string `xml:"HOTELRESTAURANT,attr" json:"hotelrestaurant,omitempty"`
					BAR             string `xml:"BAR,attr" json:"bar,omitempty"`
				} `xml:"angeschl_gastronomie" json:"angeschl_gastronomie,omitempty"`
				Brauereibindung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"brauereibindung" json:"brauereibindung,omitempty"`
				Sporteinrichtungen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"sporteinrichtungen" json:"sporteinrichtungen,omitempty"`
				Wellnessbereich struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wellnessbereich" json:"wellnessbereich,omitempty"`
				Serviceleistungen struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					BETREUTESWOHNEN string `xml:"BETREUTES_WOHNEN,attr" json:"betreutes_wohnen,omitempty"`
					CATERING        string `xml:"CATERING,attr" json:"catering,omitempty"`
					REINIGUNG       string `xml:"REINIGUNG,attr" json:"reinigung,omitempty"`
					EINKAUF         string `xml:"EINKAUF,attr" json:"einkauf,omitempty"`
					WACHDIENST      string `xml:"WACHDIENST,attr" json:"wachdienst,omitempty"`
				} `xml:"serviceleistungen" json:"serviceleistungen,omitempty"`
				TelefonFerienimmobilie struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"telefon_ferienimmobilie" json:"telefon_ferienimmobilie,omitempty"`
				Sicherheitstechnik struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					ALARMANLAGE string `xml:"ALARMANLAGE,attr" json:"alarmanlage,omitempty"`
					KAMERA      string `xml:"KAMERA,attr" json:"kamera,omitempty"`
					POLIZEIRUF  string `xml:"POLIZEIRUF,attr" json:"polizeiruf,omitempty"`
				} `xml:"sicherheitstechnik" json:"sicherheitstechnik,omitempty"`
				Unterkellert struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					Keller string `xml:"keller,attr" json:"keller,omitempty"`
				} `xml:"unterkellert" json:"unterkellert,omitempty"`
				Abstellraum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"abstellraum" json:"abstellraum,omitempty"`
				Fahrradraum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"fahrradraum" json:"fahrradraum,omitempty"`
				Rolladen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"rolladen" json:"rolladen,omitempty"`
				Dachform struct {
					Text             string `xml:",chardata" json:"text,omitempty"`
					KRUEPPELWALMDACH string `xml:"KRUEPPELWALMDACH,attr" json:"krueppelwalmdach,omitempty"`
					MANSARDDACH      string `xml:"MANSARDDACH,attr" json:"mansarddach,omitempty"`
					PULTDACH         string `xml:"PULTDACH,attr" json:"pultdach,omitempty"`
					SATTELDACH       string `xml:"SATTELDACH,attr" json:"satteldach,omitempty"`
					WALMDACH         string `xml:"WALMDACH,attr" json:"walmdach,omitempty"`
					FLACHDACH        string `xml:"FLACHDACH,attr" json:"flachdach,omitempty"`
					PYRAMIDENDACH    string `xml:"PYRAMIDENDACH,attr" json:"pyramidendach,omitempty"`
				} `xml:"dachform" json:"dachform,omitempty"`
				Bauweise struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					MASSIV      string `xml:"MASSIV,attr" json:"massiv,omitempty"`
					FERTIGTEILE string `xml:"FERTIGTEILE,attr" json:"fertigteile,omitempty"`
					HOLZ        string `xml:"HOLZ,attr" json:"holz,omitempty"`
				} `xml:"bauweise" json:"bauweise,omitempty"`
				Ausbaustufe struct {
					Text                            string `xml:",chardata" json:"text,omitempty"`
					BAUSATZHAUS                     string `xml:"BAUSATZHAUS,attr" json:"bausatzhaus,omitempty"`
					AUSBAUHAUS                      string `xml:"AUSBAUHAUS,attr" json:"ausbauhaus,omitempty"`
					SCHLUESSELFERTIGMITKELLER       string `xml:"SCHLUESSELFERTIGMITKELLER,attr" json:"schluesselfertigmitkeller,omitempty"`
					SCHLUESSELFERTIGOHNEBODENPLATTE string `xml:"SCHLUESSELFERTIGOHNEBODENPLATTE,attr" json:"schluesselfertigohnebodenplatte,omitempty"`
					SCHLUESSELFERTIGMITBODENPLATTE  string `xml:"SCHLUESSELFERTIGMITBODENPLATTE,attr" json:"schluesselfertigmitbodenplatte,omitempty"`
				} `xml:"ausbaustufe" json:"ausbaustufe,omitempty"`
				Energietyp struct {
					Text           string `xml:",chardata" json:"text,omitempty"`
					PASSIVHAUS     string `xml:"PASSIVHAUS,attr" json:"passivhaus,omitempty"`
					NIEDRIGENERGIE string `xml:"NIEDRIGENERGIE,attr" json:"niedrigenergie,omitempty"`
					NEUBAUSTANDARD string `xml:"NEUBAUSTANDARD,attr" json:"neubaustandard,omitempty"`
					KFW40          string `xml:"KFW40,attr" json:"kfw40,omitempty"`
					KFW60          string `xml:"KFW60,attr" json:"kfw60,omitempty"`
				} `xml:"energietyp" json:"energietyp,omitempty"`
				Bibliothek struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bibliothek" json:"bibliothek,omitempty"`
				Dachboden struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"dachboden" json:"dachboden,omitempty"`
				Gaestewc struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gaestewc" json:"gaestewc,omitempty"`
				Kabelkanaele struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kabelkanaele" json:"kabelkanaele,omitempty"`
				Seniorengerecht struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"seniorengerecht" json:"seniorengerecht,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Feld struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"name" json:"name,omitempty"`
						Wert struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"wert" json:"wert,omitempty"`
						Typ struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"typ" json:"typ,omitempty"`
						Modus struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"modus" json:"modus,omitempty"`
					} `xml:"feld" json:"feld,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"ausstattung" json:"ausstattung,omitempty"`
			ZustandAngaben struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Baujahr struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"baujahr" json:"baujahr,omitempty"`
				Letztemodernisierung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"letztemodernisierung" json:"letztemodernisierung,omitempty"`
				Zustand struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					ZustandArt string `xml:"zustand_art,attr" json:"zustand_art,omitempty"`
				} `xml:"zustand" json:"zustand,omitempty"`
				Alter struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					AlterAttr string `xml:"alter_attr,attr" json:"alter_attr,omitempty"`
				} `xml:"alter" json:"alter,omitempty"`
				BebaubarNach struct {
					Text         string `xml:",chardata" json:"text,omitempty"`
					BebaubarAttr string `xml:"bebaubar_attr,attr" json:"bebaubar_attr,omitempty"`
				} `xml:"bebaubar_nach" json:"bebaubar_nach,omitempty"`
				Erschliessung struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					ErschlAttr string `xml:"erschl_attr,attr" json:"erschl_attr,omitempty"`
				} `xml:"erschliessung" json:"erschliessung,omitempty"`
				ErschliessungUmfang struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"erschliessung_umfang" json:"erschliessung_umfang,omitempty"`
				Bauzone struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bauzone" json:"bauzone,omitempty"`
				Altlasten struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"altlasten" json:"altlasten,omitempty"`
				Energiepass struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Epart struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"epart" json:"epart,omitempty"`
					GueltigBis struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"gueltig_bis" json:"gueltig_bis,omitempty"`
					Energieverbrauchkennwert struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"energieverbrauchkennwert" json:"energieverbrauchkennwert,omitempty"`
					Mitwarmwasser struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"mitwarmwasser" json:"mitwarmwasser,omitempty"`
					Endenergiebedarf struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"endenergiebedarf" json:"endenergiebedarf,omitempty"`
					Primaerenergietraeger struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"primaerenergietraeger" json:"primaerenergietraeger,omitempty"`
					Stromwert struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"stromwert" json:"stromwert,omitempty"`
					Waermewert struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"waermewert" json:"waermewert,omitempty"`
					Wertklasse struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"wertklasse" json:"wertklasse,omitempty"`
					Ausstelldatum struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"ausstelldatum" json:"ausstelldatum,omitempty"`
					Jahrgang struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"jahrgang" json:"jahrgang,omitempty"`
					Gebaeudeart struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"gebaeudeart" json:"gebaeudeart,omitempty"`
				} `xml:"energiepass" json:"energiepass,omitempty"`
				Verkaufstatus struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"verkaufstatus" json:"verkaufstatus,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"zustand_angaben" json:"zustand_angaben,omitempty"`
			Bewertung struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Feld struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Name struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"name" json:"name,omitempty"`
					Wert struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"wert" json:"wert,omitempty"`
					Typ struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"typ" json:"typ,omitempty"`
					Modus struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"modus" json:"modus,omitempty"`
				} `xml:"feld" json:"feld,omitempty"`
			} `xml:"bewertung" json:"bewertung,omitempty"`
			Infrastruktur struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				Zulieferung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zulieferung" json:"zulieferung,omitempty"`
				Ausblick struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Blick string `xml:"blick,attr" json:"blick,omitempty"`
				} `xml:"ausblick" json:"ausblick,omitempty"`
				Distanzen struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					DistanzZu string `xml:"distanz_zu,attr" json:"distanz_zu,omitempty"`
				} `xml:"distanzen" json:"distanzen,omitempty"`
				DistanzenSport struct {
					Text           string `xml:",chardata" json:"text,omitempty"`
					DistanzZuSport string `xml:"distanz_zu_sport,attr" json:"distanz_zu_sport,omitempty"`
				} `xml:"distanzen_sport" json:"distanzen_sport,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"infrastruktur" json:"infrastruktur,omitempty"`
			Freitexte struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				Objekttitel struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"objekttitel" json:"objekttitel,omitempty"`
				Dreizeiler struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"dreizeiler" json:"dreizeiler,omitempty"`
				Lage struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"lage" json:"lage,omitempty"`
				AusstattBeschr struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"ausstatt_beschr" json:"ausstatt_beschr,omitempty"`
				Objektbeschreibung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"objektbeschreibung" json:"objektbeschreibung,omitempty"`
				SonstigeAngaben struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"sonstige_angaben" json:"sonstige_angaben,omitempty"`
				ObjektText struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Lang string `xml:"lang,attr" json:"lang,omitempty"`
				} `xml:"objekt_text" json:"objekt_text,omitempty"`
			} `xml:"freitexte" json:"freitexte,omitempty"`
			Anhaenge struct {
				Text   string `xml:",chardata" json:"text,omitempty"`
				Anhang struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					Location    string `xml:"location,attr" json:"location,omitempty"`
					Gruppe      string `xml:"gruppe,attr" json:"gruppe,omitempty"`
					Anhangtitel struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"anhangtitel" json:"anhangtitel,omitempty"`
					Format struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"format" json:"format,omitempty"`
					Check struct {
						Text  string `xml:",chardata" json:"text,omitempty"`
						Ctype string `xml:"ctype,attr" json:"ctype,omitempty"`
					} `xml:"check" json:"check,omitempty"`
					Daten struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Pfad struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"pfad" json:"pfad,omitempty"`
					} `xml:"daten" json:"daten,omitempty"`
				} `xml:"anhang" json:"anhang,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
			} `xml:"anhaenge" json:"anhaenge,omitempty"`
			VerwaltungObjekt struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				ObjektadresseFreigeben struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"objektadresse_freigeben" json:"objektadresse_freigeben,omitempty"`
				VerfuegbarAb struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"verfuegbar_ab" json:"verfuegbar_ab,omitempty"`
				Abdatum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"abdatum" json:"abdatum,omitempty"`
				Bisdatum struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"bisdatum" json:"bisdatum,omitempty"`
				MinMietdauer struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					MinDauer string `xml:"min_dauer,attr" json:"min_dauer,omitempty"`
				} `xml:"min_mietdauer" json:"min_mietdauer,omitempty"`
				MaxMietdauer struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					MaxDauer string `xml:"max_dauer,attr" json:"max_dauer,omitempty"`
				} `xml:"max_mietdauer" json:"max_mietdauer,omitempty"`
				Versteigerungstermin struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"versteigerungstermin" json:"versteigerungstermin,omitempty"`
				WbsSozialwohnung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"wbs_sozialwohnung" json:"wbs_sozialwohnung,omitempty"`
				Vermietet struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"vermietet" json:"vermietet,omitempty"`
				Gruppennummer struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gruppennummer" json:"gruppennummer,omitempty"`
				Zugang struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"zugang" json:"zugang,omitempty"`
				Laufzeit struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"laufzeit" json:"laufzeit,omitempty"`
				MaxPersonen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"max_personen" json:"max_personen,omitempty"`
				Nichtraucher struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"nichtraucher" json:"nichtraucher,omitempty"`
				Haustiere struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"haustiere" json:"haustiere,omitempty"`
				Geschlecht struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					GeschlAttr string `xml:"geschl_attr,attr" json:"geschl_attr,omitempty"`
				} `xml:"geschlecht" json:"geschlecht,omitempty"`
				Denkmalgeschuetzt struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"denkmalgeschuetzt" json:"denkmalgeschuetzt,omitempty"`
				AlsFerien struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"als_ferien" json:"als_ferien,omitempty"`
				GewerblicheNutzung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gewerbliche_nutzung" json:"gewerbliche_nutzung,omitempty"`
				Branchen struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"branchen" json:"branchen,omitempty"`
				Hochhaus struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"hochhaus" json:"hochhaus,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
			} `xml:"verwaltung_objekt" json:"verwaltung_objekt,omitempty"`
			VerwaltungTechn struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				ObjektnrIntern struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"objektnr_intern" json:"objektnr_intern,omitempty"`
				ObjektnrExtern struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"objektnr_extern" json:"objektnr_extern,omitempty"`
				Aktion struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					Aktionart string `xml:"aktionart,attr" json:"aktionart,omitempty"`
				} `xml:"aktion" json:"aktion,omitempty"`
				AktivVon struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"aktiv_von" json:"aktiv_von,omitempty"`
				AktivBis struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"aktiv_bis" json:"aktiv_bis,omitempty"`
				OpenimmoObid struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"openimmo_obid" json:"openimmo_obid,omitempty"`
				KennungUrsprung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"kennung_ursprung" json:"kennung_ursprung,omitempty"`
				StandVom struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"stand_vom" json:"stand_vom,omitempty"`
				WeitergabeGenerell struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"weitergabe_generell" json:"weitergabe_generell,omitempty"`
				WeitergabePositiv struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"weitergabe_positiv" json:"weitergabe_positiv,omitempty"`
				WeitergabeNegativ struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"weitergabe_negativ" json:"weitergabe_negativ,omitempty"`
				GruppenKennung struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"gruppen_kennung" json:"gruppen_kennung,omitempty"`
				Master struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					Visible string `xml:"visible,attr" json:"visible,omitempty"`
				} `xml:"master" json:"master,omitempty"`
				Sprache struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"sprache" json:"sprache,omitempty"`
				UserDefinedSimplefield struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
				} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
				UserDefinedAnyfield struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
				UserDefinedExtend struct {
					Text string `xml:",chardata" json:"text,omitempty"`
				} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
			} `xml:"verwaltung_techn" json:"verwaltung_techn,omitempty"`
			UserDefinedSimplefield struct {
				Text     string `xml:",chardata" json:"text,omitempty"`
				Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
			} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
			UserDefinedAnyfield struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Plaintext []struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Fieldname  string `xml:"fieldname,attr" json:"fieldname,omitempty"`
					Fieldvalue string `xml:"fieldvalue,attr" json:"fieldvalue,omitempty"`
				} `xml:"plaintext" json:"plaintext,omitempty"`
				Flag []struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Fieldname  string `xml:"fieldname,attr" json:"fieldname,omitempty"`
					Fieldvalue string `xml:"fieldvalue,attr" json:"fieldvalue,omitempty"`
				} `xml:"flag" json:"flag,omitempty"`
			} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
			UserDefinedExtend struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"user_defined_extend" json:"user_defined_extend,omitempty"`
		} `xml:"immobilie" json:"immobilie,omitempty"`
		Impressum struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"impressum" json:"impressum,omitempty"`
		ImpressumStrukt struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			Firmenname struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"firmenname" json:"firmenname,omitempty"`
			Firmenanschrift struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"firmenanschrift" json:"firmenanschrift,omitempty"`
			Telefon struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"telefon" json:"telefon,omitempty"`
			Vertretungsberechtigter struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"vertretungsberechtigter" json:"vertretungsberechtigter,omitempty"`
			Berufsaufsichtsbehoerde struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"berufsaufsichtsbehoerde" json:"berufsaufsichtsbehoerde,omitempty"`
			Handelsregister struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"handelsregister" json:"handelsregister,omitempty"`
			HandelsregisterNr struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"handelsregister_nr" json:"handelsregister_nr,omitempty"`
			UmsstID struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"umsst-id" json:"umsst-id,omitempty"`
			Steuernummer struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"steuernummer" json:"steuernummer,omitempty"`
			Weiteres struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"weiteres" json:"weiteres,omitempty"`
		} `xml:"impressum_strukt" json:"impressum_strukt,omitempty"`
		UserDefinedSimplefield struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
		} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
		UserDefinedAnyfield struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
	} `xml:"anbieter" json:"anbieter,omitempty"`
	UserDefinedSimplefield struct {
		Text     string `xml:",chardata" json:"text,omitempty"`
		Feldname string `xml:"feldname,attr" json:"feldname,omitempty"`
	} `xml:"user_defined_simplefield" json:"user_defined_simplefield,omitempty"`
	UserDefinedAnyfield struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"user_defined_anyfield" json:"user_defined_anyfield,omitempty"`
} 

