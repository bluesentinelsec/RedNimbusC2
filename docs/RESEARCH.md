# RESEARCH

## Why use serverless C2?
- simplifies C2 infrastructure deployment and maintenance
  - no servers to manage (accounts, patching, hardening, etc.)
  - no need to manage TLS certs or domain names

- it is hard for defenders to to detect serverless C2
  - C2 rides trusted infrastructure
  - traffic is encrypted (can TLS inspection mitigate this?)
  - difficult or impractical to block cloud provider traffic

- red team accountability
  - red team access via IAM roles
  - CloudWatch for capturing red operator activity

## What does serverless C2 look like today?
- potentially one case observed in the wild from Arista / Awake Security
- one public tool by Rob Goyette (hackerrob) - https://github.com/hackerob/ServerlessC2
- several examples of using serverless for Cobalt Strike Beacon redirection


## References

1. https://www.arista.com/assets/data/pdf/CaseStudies/Case-Study-Serverless-C2-Cloud.pdf
  - June 28, 2022
  - Arista NDR detected serverless C2 in Azure cloud
  - claims behavior was detected via anomalous network behavior originating from a MS-Office add-in
  - not really clear how they know this to be serverless C2
  - "Additionally, the C2 server was serverless code in the Azure cloud, so all that was visible on the network was an encrypted tunnel to a subdomain of azurewebsites[.]net."

2. https://awakesecurity.com/wp-content/uploads/2020/01/CS-Financial-Services-Serverless-C2-Cloud.pdf
  - this is a duplicitive report of the Arista 1 pager
  - dated 2022? ARISTA acquired Awake Security at some point

3. https://securityboulevard.com/2019/07/detecting-the-impossible-serverless-c2-in-the-cloud/
  - Author Gary Golomb July 2019
  - this seems to describe the same activty in the Arista+Awake article
  - the C2 server was using TLS (this is expected)
  - only evidence of serverless C2 seems to be the encrypted tunnel to azurewebsites.net?
  - describes an attack using Office add-ins, which are DLLs renamed to .wll and placed in %appdata%RoamingMicrosoftWordstartup

4. https://aristanetworks.force.com/AristaCommunity/s/article/Threat-Hunting-Series-Detecting-Command-Control-in-the-Cloud
  - another Arista article; has some prose describing the challenges posed by serverless C2
  - "This is terrifying from a threat detection and hunting perspective because the vast majority of a company’s Internet traffic is already going to Microsoft, Google, Amazon, and Cloudflare – and all of it is pretty much encrypted, too. When running this way, the C2 traffic has the same hosting, certificate, and server characteristics as the vast majority of traffic to/from most enterprises. As our customer pointed out to us, they have yet to find a Network Traffic Analysis (NTA) solution that can handle a situation like this."

5. https://howto.thec2matrix.com/attack-infrastructure/redirectors
  - this provides several relevant URLs to Lambda or serverless articles

6. https://blog.xpnsec.com/aws-lambda-redirector/
  - Adam Chester, TrustedSec, 2020-02-25
  - shows how to use Lambda to redirect cobalt strike beacon callbacks

7. https://rhinosecuritylabs.com/aws/hiding-cloudcobalt-strike-beacon-c2-using-amazon-apis/
  - Benjamin Caudill, Rhino Security Labs
  - no publication date listed, but way back machine lists the date as February 15th 2018 and the author as Dwight Hohnstein
  - this article actually provides a good PoC for end-to-end serverless C2
  - however, it is built around Cobalt Strike meaning you still have to manage the Team Server and the license cost

8. https://hstechdocs.helpsystems.com/manuals/cobaltstrike/current/userguide/content/externalc2spec.pdf
  - External C2 specification
  - Does AWS Lambda negate the the need for redirectors?

9. https://www.trustedsec.com/blog/front-validate-and-redirect/
  - melvin langvik February 16, 2021
  - shows a good diagram of an Azure C2 relay

10. https://redteamer.tips/servers-are-overrated-bypassing-corporate-proxies-abusing-serverless-for-fun-and-profit/
  - Jean Maes October 21 2021
  - "A lot of technology is moving/moved to the cloud. It is hard for companies to block access to these cloud services completely. By leveraging the cloud ourselves, in this case using serverless, we have a good chance of bypassing corporate outbound internet restrictions."

11. https://khast3x.club/posts/2020-02-14-Intro-Modern-Routing-Traefik-Metasploit-Docker/

12. https://github.com/hackerob/ServerlessC2
  - this project seems to do what I'm proposing - nice!
  - created in March based on commit history
  - has a web app UI

13. https://www.setsolutions.com/hiding-in-plain-sight/
  - "Just as “serverless” infrastructure is the new and emerging trend in Information Technology, it is also the new and emerging trend in malware."
  - Justin Hutchens Set Solutions August 4th 2020

14. https://www.oreilly.com/library/view/hands-on-red-team/9781788995238/55ecbb0e-1ba0-4343-b239-a7f4d1176d3f.xhtml
  - cloud-based file sharing






